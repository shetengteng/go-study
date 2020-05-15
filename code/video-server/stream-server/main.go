package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 定义中间件
type middleWareHandler struct {
	// 对router进行包裹
	r *httprouter.Router
	// 对router使用限流器
	l *ConnLimiter
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !m.l.RequireToken() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	// 执行请求
	m.r.ServeHTTP(w, req)
	// 释放请求
	m.l.ReleaseToken()
}

func CreateMiddleWareHandler(r *httprouter.Router, count int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = CreateLimiter(count)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// 下载
	router.GET("/videos/:vid-id", downloadHandler)
	// 上传
	router.POST("/upload/:vid-id", uploadHandler)
	// 测试页面
	router.GET("/testpage", testPageHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	m := CreateMiddleWareHandler(r, 2)
	http.ListenAndServe(":8001", m)
}
