package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	routers *router
}

func New() *Engine {
	return &Engine{routers: newRouter()}
}

func (e *Engine) addRouter(method, path string, handler HandlerFunc) {
	e.routers.addRoute(method, path, handler)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRouter("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRouter("POST", path, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.routers.handler(c)

}
