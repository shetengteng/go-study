package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/param", func(context *gin.Context) {
		// 如果url中的参数不存在，则使用默认值
		name := context.DefaultQuery("name", "stt")
		// url中参数不存在则返回空
		age := context.Query("age")
		context.String(200, "name is"+name+", age is"+age)
	})
	r.Run(":8089")
}
