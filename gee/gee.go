package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouteGroup
	routers *router
	groups  []*RouteGroup
}

type RouteGroup struct {
	prefix      string
	engine      *Engine
	middlewares []HandlerFunc
	parent      *RouteGroup
}

func New() *Engine {
	engine := &Engine{routers: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (g *RouteGroup) Group(prefix string) *RouteGroup {
	engine := g.engine
	newGroup := &RouteGroup{
		prefix: g.prefix + prefix,
		engine: engine,
		parent: g,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouteGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.routers.addRoute(method, pattern, handler)
}

func (group *RouteGroup) GET(path string, handler HandlerFunc) {
	group.addRoute("GET", path, handler)
}

func (group *RouteGroup) POST(path string, handler HandlerFunc) {
	group.addRoute("POST", path, handler)
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
