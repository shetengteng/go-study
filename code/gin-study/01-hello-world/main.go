package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	// 创建路由
	router := gin.Default()
	// 绑定路由规则, gin.Context 封装了 request 和 response
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"helloworld")
	})
	// 监听端口
	router.Run(":8089")
}
