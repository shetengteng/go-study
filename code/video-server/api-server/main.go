package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// 创建用户
	router.POST("/user", CreateUser)
	// 用户登录
	router.POST("/user/:username", Login)

	return router
}

func main() {
	router := RegisterHandlers()
	// 添加过滤器，和过滤的方法
	// routerHasFilter := AddFilter(router, addUsernameInSession)

	routerHasFilter := WrapFilter(router).
		AddFilterMethod(addUsernameInSession).
		GetHandler()

	// 监听
	http.ListenAndServe(":8000", routerHasFilter)
}

// 原理
// listen-->RegisterHandlers-->handlers
// 每个请求的handlers是依据不同的goroutine创建的

// api 的主要分层模块
// handler->validation{1.request是否合法，2.user是否合法}->business logic-> response
// validation
// 		1.data model
// 		2.error handling

// 安装  xxx\ go install
// 调试 xxx\go run .
