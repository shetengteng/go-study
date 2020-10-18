package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 本地中间件，针对某个请求进行处理
func main() {
	r := gin.Default()
	// 针对该请求使用中间件
	r.GET("/middleware2", MiddlewareHandler(), func(context *gin.Context) {
		// 取值
		val, _ := context.Get("my-key")
		val2, _ := context.Get("my-key2")
		fmt.Println("-----get val", val)
		context.JSON(200, gin.H{"value": val, "value2": val2})
	})
	r.Run(":8090")
}

func MiddlewareHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		fmt.Println("middleware start")
		context.Set("my-key", "middleware")
		// 执行请求本地操作，操作执行成功后计算耗时
		context.Next()
		fmt.Println("middleware finish", context.Writer.Status())
		fmt.Println("time:", time.Since(now))
	}
}
