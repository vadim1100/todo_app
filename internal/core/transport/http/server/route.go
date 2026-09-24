package core_http_server

import "net/http"

type Route struct {
	Method string
	Path string
	Handler http.HandlerFunc
	Protected bool
}

func NewRoute(
	method string,
	path string,
	handler http.HandlerFunc,
	protected bool,
) Route {
	return Route{
		Method: method,
		Path: path,
		Handler: handler,
		Protected: protected,
	}
}