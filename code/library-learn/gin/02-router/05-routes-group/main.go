package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// httproter会将所有路由规则构造一颗前缀树
func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/name", name)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/name", name)
	}

	r.Run(":8089")
}

func login(context *gin.Context) {
	name := context.DefaultQuery("name", "stt")
	context.String(200, fmt.Sprintf("name is %s", name))
}

func name(context *gin.Context) {
	name := context.DefaultQuery("name", "stt1")
	context.String(200, fmt.Sprintf("name is %s", name))
}
