package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//表单传输为post请求，http常见的传输格式为四种：
//application/json
//application/x-www-form-urlencoded
//application/xml
//multipart/form-data
//表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

func main() {
	r := gin.Default()
	r.POST("/form", func(context *gin.Context) {
		age := context.DefaultPostForm("age","22")
		name := context.PostForm("username")
		passwd := context.PostForm("passwd")
		context.String(200,fmt.Sprintf("username:%s passwd %s age %s",name,passwd,age))
	})
	r.Run(":8089")
}
