package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type user struct {
	// age > 10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/validator", func(context *gin.Context) {
		var user user
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(500, gin.H{"error": fmt.Sprint(err)})
			return
		}
		//1. %v    只输出所有的值
		//2. %+v 先输出字段类型，再输出该字段的值
		//3. %#v 先输出结构体名字值，再输出结构体（字段类型+字段的值）
		fmt.Println(user.Birthday.String())
		context.JSON(200, gin.H{"user": fmt.Sprintf("%#v", user)})
	})
	r.Run(":8090")
}
