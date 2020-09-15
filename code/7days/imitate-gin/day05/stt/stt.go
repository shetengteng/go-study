package stt

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 中间件
	parent      *RouterGroup  // 支持嵌套，当前group的父group
	engine      *Engine       // 创建时，全局唯一实例
}

// engine作为顶层分组
type Engine struct {
	*RouterGroup // RouterGroup对象，对外暴露
	router       *router
	groups       []*RouterGroup // 存储所有的router group
}

// 构造器
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 创建group
// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix, // 每个group上是前面group的url前缀之和
		engine: engine,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 使用了(group *RouterGroup) 但是 engine也可以调用该方法，继承了RouterGroup的方法
func (group *RouterGroup) addRoute(method string, comp string, handlerFunc HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handlerFunc)
}

func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("GET", pattern, handlerFunc)
}

func (group *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("POST", pattern, handlerFunc)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) Run(addr string) (err error) {
	// engine 实现了 ServeHTTP方法，等价于实现了http.Handler接口
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		// 如果url匹配到了相应的group的前缀，那么就走相应group内部的中间件
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context := newContext(w, r)
	// 将中间件放入到handlers中
	context.handlers = middlewares
	engine.router.handle(context)
}
