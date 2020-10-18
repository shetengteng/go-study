package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(context *gin.Context) {
		cookie, err := context.Cookie("my-cookie")
		if err != nil {
			cookie = "noset"
			// maxAge 单位s
			// path cookie 所在的目录
			// domain 域名
			// secure 是否智能通过https访问
			// httpOnly 是否允许别人通过js获取自己的cookie
			context.SetCookie("my-cookie", "stt", 60, "/",
				"localhost", false, true)
		}
		fmt.Println("cookie:", cookie)
		context.JSON(200, gin.H{"cookie": cookie})
	})
	r.Run(":8090")
}
