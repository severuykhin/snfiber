package snfiber

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes []*Route
}

type path string
type groupPrefix path

type Route struct {
	Method      string
	Path        path
	Handler     requestHandler
	Middlewares []requestHandler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Post(path path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodPost, path, handler, middlewares...)
}

func (r *Router) Get(path path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodGet, path, handler, middlewares...)
}

func (r *Router) Patch(path path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodPatch, path, handler, middlewares...)
}

func (r *Router) Delete(path path, handler requestHandler, middlewares ...requestHandler) *Route {
	return r.Add(http.MethodDelete, path, handler, middlewares...)
}

func (r *Router) Add(method string, path path, handler requestHandler, middlewares ...requestHandler) *Route {
	route := Route{
		Method:      method,
		Path:        path,
		Handler:     handler,
		Middlewares: middlewares,
	}

	r.routes = append(r.routes, &route)
	return &route
}

func (r *Router) Group(prefix groupPrefix, routes ...*Route) {
	for _, route := range routes {
		route.Path = path(fmt.Sprintf("%s%s", prefix, route.Path))
	}
}
