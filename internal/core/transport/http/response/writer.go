package core_http_response

import "net/http"

var (
	StatusCodeUninitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(rw http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: rw,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) GetStatusCode() int {
	if rw.statusCode == StatusCodeUninitialized {
		panic("no status code set")
	}
	return rw.statusCode
}