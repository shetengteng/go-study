package stt

import (
	"net/http"
)

type HandlerFunc func(c *Context)

// 通过key: method-url 匹配到相应的路由方法进行执行
type Engine struct {
	router *router
}

// 构造器
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 向map中放入路由处理逻辑
func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router.addRoute(method, pattern, handlerFunc)
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.addRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.addRoute("POST", pattern, handlerFunc)
}

func (engine *Engine) Run(addr string) (err error) {
	// engine 实现了 ServeHTTP方法，等价于实现了http.Handler接口
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	engine.router.handle(newContext(w, r))
}
