package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	routers map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{routers: make(map[string]HandlerFunc)}
}

func (e *Engine) addRouter(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	e.routers[key] = handler
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	key := "GET" + "-" + path
	e.routers[key] = handler
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	key := "POST" + "-" + path
	e.routers[key] = handler
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.routers[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintln(w, "404 NOT FOUND")
	}
}
