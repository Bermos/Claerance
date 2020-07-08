package routing

import "net/http"

type Route struct {
	pattern string
	handler Handler
}

type Handler func(w http.ResponseWriter, r http.Request, p map[string]string)

func NewRoute(pattern string, handler Handler) Route {

	return Route{pattern: pattern, handler: handler}
}

func (r Route) match(url string) (Handler, map[string]string) {
	return r.handler, map[string]string{}
}
