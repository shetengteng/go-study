package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// main-->middleware 做一些校验，流控处理等 (filter)--> defs(message,err) --> handlers --> ops --> response

// 实现Handler的ServeHTTP(ResponseWriter, *Request) 接口
type filter struct {
	router  *httprouter.Router
	methods []func(r *http.Request) bool
}

// 实现了ServeHTTP 方法，那么filter就是http.Handler对象
func (f *filter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 每一个http请求都会检查一下session
	for _, m := range f.methods {
		if !m(req) {
			return
		}
	}
	f.router.ServeHTTP(w, req)
}

// AddFilter 添加过滤方法，如果重复调用，返回新的filter对象
func AddFilter(r *httprouter.Router, m ...func(r *http.Request) bool) http.Handler {
	filter := filter{}
	filter.router = r
	filter.methods = make([]func(r *http.Request) bool, 0)
	filter.methods = append(filter.methods, m...)
	return &filter
}

// WrapFilter 包裹一个过滤器
func WrapFilter(r *httprouter.Router) *filter {
	f := filter{}
	f.router = r
	f.methods = make([]func(r *http.Request) bool, 0)
	return &f
}

func (f *filter) AddFilterMethod(m ...func(r *http.Request) bool) *filter {
	f.methods = append(f.methods, m...)
	return f
}

func (f *filter) GetHandler() http.Handler {
	return f
}
