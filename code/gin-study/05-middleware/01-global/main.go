package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 所有请求都经过此中间件，可以看做拦截器
// 添加next方法，则是aop中的环绕通知，没有next方法则是aop中的前置通知
func main() {
	r := gin.Default()
	//gin 默认使用了2个中间件Logger(), Recovery()
	// 所有请求都使用该中间件
	r.Use(MiddlewareHandler())
	r.Use(MiddlewareHandler2())
	r.GET("/middleware", func(context *gin.Context) {
		// 取值
		val, _ := context.Get("my-key")
		val2, _ := context.Get("my-key2")
		fmt.Println("-----get val", val)
		context.JSON(200, gin.H{"value": val, "value2": val2})
	})
	r.Run(":8090")

	// 执行结果
	//	middleware start
	//	middleware2 start # 由于没有在middleware中添加next方法，因此没有包裹的现象
	//	middleware2 finish
	//	time2: 0s
	//	-----get val middleware
	//	middleware finish 200
	//	time: 995.8µs

}

// 可以看做为一个拦截器，进行设置参数处理
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

func MiddlewareHandler2() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		fmt.Println("middleware2 start")
		context.Set("my-key2", "middleware2")
		fmt.Println("middleware2 finish")
		fmt.Println("time2:", time.Since(now))
	}
}
