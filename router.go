package snfiber

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes []*Route
}

type Path string

type Route struct {
	Method      string
	Path        Path
	Handler     requestHandler
	Middlewares []requestHandler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Post(path Path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodPost, path, handler, middlewares...)
}

func (r *Router) Get(path Path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodGet, path, handler, middlewares...)
}

func (r *Router) Patch(path Path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodPatch, path, handler, middlewares...)
}

func (r *Router) Delete(path Path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodDelete, path, handler, middlewares...)
}

func (r *Router) Add(method string, path Path, handler requestHandler, middlewares ...requestHandler) *Route {
	route := Route{
		Method:      method,
		Path:        path,
		Handler:     handler,
		Middlewares: middlewares,
	}

	r.routes = append(r.routes, &route)
	return &route
}

func (r *Router) Group(prefix Path, routes ...*Route) {
	for _, route := range routes {
		route.Path = Path(fmt.Sprintf("%s%s", prefix, route.Path))
	}
}
