package stt

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// 通过key: method-url 匹配到相应的路由方法进行执行
type Engine struct {
	router map[string]HandlerFunc
}

// 构造器
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 向map中放入路由处理逻辑
func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handlerFunc
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
	// 解析方法中的url和methodName，进行engine中router的匹配
	key := r.Method + "-" + r.URL.Path
	if handlerFunc, ok := engine.router[key]; ok {
		handlerFunc(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %q\n", r.URL)
	}
}
