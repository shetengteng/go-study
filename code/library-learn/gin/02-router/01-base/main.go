package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//gin 框架中采用的路由库是基于httprouter做的
//地址为：https://github.com/julienschmidt/httprouter

func main() {
	r := gin.Default()
	r.GET("/base", func(context *gin.Context) {
		context.String(http.StatusOK,"hello world")
	})
	r.POST("/base", func(context *gin.Context) {
		fmt.Println(context)
	})
	r.PUT("/base")
	r.Run(":8089")
}
