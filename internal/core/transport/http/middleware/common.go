package core_http_middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_service_jwt "github.com/vadim1100/todo_app/internal/core/service/jwt"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-ID"
const authorizationHeader = "Authorization"

type authKey struct{}

var userIDKey authKey

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			
			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), core_logger.LogKey, l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"a panic occurred when the http handler attempted to execute",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()

			logger.Debug(
				">>>>> incoming http request",
				zap.Time("time", time.Now().UTC()),
			)

			next.ServeHTTP(rw, r)

			logger.Debug(
				"<<<<< done http request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Since(before)),
			)
		})
	}
}

func Auth(jwtManager *core_service_jwt.Manager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get(authorizationHeader)
			logger := core_logger.FromContext(r.Context())
			responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

			if authHeader == "" {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"missing authorization header",
				)
				return
			}

			scheme, token, ok := strings.Cut(authHeader, " ")
			if !ok || scheme != "Bearer" || token == "" {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"invalid authorization header",
				)
				return 
			}

			claims, err := jwtManager.Parse(token)
			if err != nil {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"invalid token",
				)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UserIDFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIDKey).(int)
	return id, ok
}