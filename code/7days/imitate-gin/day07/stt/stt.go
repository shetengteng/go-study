package stt

import (
	"html/template"
	"net/http"
	"path"
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

	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

// 构造器
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Recovery())
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
	context.engine = engine
	// 将中间件放入到handlers中
	context.handlers = middlewares
	engine.router.handle(context)
}

// 静态文件处理
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	// 获得到资源的绝对路径
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		// 从路径url中获取文件名称 要求 配置为 /xx/*filepath 表示获取xx下的filepath所有资源
		file := c.Param("filepath")
		// 判断文件是否存在
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

// 设置静态资源文件路径
// 用户可以将磁盘上的某个文件夹root映射到路由relativePath
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	// 获取相对路径下的所有文件
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}

// html模板处理
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(
		template.New("").
			Funcs(engine.funcMap).
			ParseGlob(pattern))
}
