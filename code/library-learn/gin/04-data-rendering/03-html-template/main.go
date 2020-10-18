package main

import "github.com/gin-gonic/gin"

//gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换
//LoadHTMLGlob()方法可以加载模板文件
// 注意需要使用go build main.go后测试，由于go run的路径不同，导致文件夹找不到
func main() {

	r := gin.Default()
	//r.LoadHTMLGlob("page/*")
	// 支持多层文件夹，前提是要有文件存在
	r.LoadHTMLGlob("page/**/*")
	// 静态文件
	//r.Static("/assets", "./assets")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{"title": "ceshi", "key": 123456})
	})
	r.Run(":8090")
}
