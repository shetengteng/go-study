package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/param/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action") // 【/xxxx】 需要去除/
		action = strings.Trim(action,"/")
		context.String(http.StatusOK, "name is "+name+" ,action is "+action)
	})
	r.Run(":8089")
}
