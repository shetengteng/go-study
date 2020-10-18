package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 访问home，如果没有cookie则返回失败，访问login获取到cookie，再访问home成功
func main() {

	r := gin.Default()
	r.GET("/home", AuthMiddleware(), func(context *gin.Context) {
		cookie, _ := context.Cookie("my-cookie")
		context.JSON(200, gin.H{"msg": "ok home", "cookie": cookie})
	})
	r.GET("/login", func(context *gin.Context) {
		context.SetCookie("my-cookie", "stt", 60, "/",
			"localhost", false, true)
		context.JSON(200, gin.H{"msg": "ok login"})
	})
	r.Run(":8090")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Cookie("my-cookie")
		if err == nil && cookie == "stt" {
			context.Next()
			return
		}
		context.JSON(http.StatusUnauthorized, gin.H{"msg": "no cookie"})
		context.Abort()
	}
}
