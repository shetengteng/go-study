package stt

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node // 用于匹配pattern，返回的pattern 在handlers中查找handlerFunc，同时返回url中的参数
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析 url, 将注册的url 解析成parts数组
func parsePattern(pattern string) []string {
	// 定义一个parts切片
	parts := make([]string, 0)
	for _, item := range strings.Split(pattern, "/") {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' { // 如果是 * 那么就不继续添加了
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handlerFunc

	_, ok := r.roots[method]
	if !ok {
		// 创建根节点
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parsePattern(pattern), 0)
}

// 返回节点，以及解析的参数，对于 * 解析有点弱
func (r *router) getRoute(method string, path string) (*node, map[string]string) {

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	searchParts := parsePattern(path)
	node := root.search(searchParts, 0)
	if node == nil {
		return nil, nil
	}

	// 解析参数
	params := make(map[string]string)
	configParts := parsePattern(node.pattern)
	for index, part := range configParts {
		if part[0] == ':' {
			// 得到search的值
			params[part[1:]] = searchParts[index]
		} else if part[0] == '*' && len(part) > 1 {
			// 将原先url中的匹配项进行合并
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}
	return node, params
}

// 为了支持中间件，对中间件内方法进行调用
func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		key := c.Method + "-" + node.pattern
		c.Params = params
		// 将此处的调用和中间件的调用整合到一起
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s \n", c.Path)
		})
	}
	// 执行
	c.Next()
}
