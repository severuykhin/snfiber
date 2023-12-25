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
	Method  string
	Path    path
	Handler requestHandler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Post(path path, handler requestHandler) *Route {
	return r.Add(http.MethodPost, path, handler)
}

func (r *Router) Get(path path, handler requestHandler) *Route {
	return r.Add(http.MethodGet, path, handler)
}

func (r *Router) Patch(path path, handler requestHandler) *Route {
	return r.Add(http.MethodPost, path, handler)
}

func (r *Router) Delete(path path, handler requestHandler) *Route {
	return r.Add(http.MethodGet, path, handler)
}

func (r *Router) Add(method string, path path, handler requestHandler) *Route {
	route := Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}

	r.routes = append(r.routes, &route)
	return &route
}

func (r *Router) Group(prefix groupPrefix, routes ...*Route) {
	for _, route := range routes {
		route.Path = path(fmt.Sprintf("%s%s", prefix, route.Path))
	}
}
