package main

import "github.com/gin-gonic/gin"

type Msg struct {
	Name string `xml:"name"`
	Age  int
}

func main() {

	r := gin.Default()
	// map转换为json返回
	r.GET("/toJson", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "ok", "status": 200})
	})
	// 结构体转为json
	r.GET("/toJson2", func(context *gin.Context) {

		// 内部定义的struct 转换为xml失败
		var msg struct {
			Name string
			Age  int
		}
		msg.Name = "stt"
		msg.Age = 22
		context.JSON(200, msg)
	})
	// map转换为xml返回,格式如下
	//<map>
	//  <msg>ok</msg>
	//  <status>200</status>
	//</map>
	r.GET("/toXML", func(context *gin.Context) {
		context.XML(200, gin.H{"msg": "ok", "status": 200})
	})

	//<Msg>
	//  <name>stt</name>
	//  <Age>22</Age>
	//</Msg>
	r.GET("/toXML2", func(context *gin.Context) {
		context.XML(200, Msg{"stt", 22})
	})

	// map转换为yaml
	r.GET("/toYaml", func(context *gin.Context) {
		context.YAML(200, gin.H{"msg": "ok", "status": 200})
	})
	r.Run(":8090")
}
