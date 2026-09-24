package core_http_server

import (
	"fmt"
	"net/http"
)

type Router struct {
	*http.ServeMux
	authMiddleware func(http.Handler) http.Handler
}

func NewRouter(authMiddleware func(http.Handler) http.Handler) *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
		authMiddleware: authMiddleware,
	}
}

func (r *Router) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		handler := http.Handler(route.Handler)
		if route.Protected {
			handler = r.authMiddleware(handler)
		}
		
		r.Handle(pattern, handler)
	}
}