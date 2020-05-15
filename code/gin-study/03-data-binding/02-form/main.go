package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/data/loginForm", func(context *gin.Context) {
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := context.Bind(&login); err != nil {
			// gin.H封装了生成json数据的工具
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": 200})
	})
	r.Run(":8089")
}
