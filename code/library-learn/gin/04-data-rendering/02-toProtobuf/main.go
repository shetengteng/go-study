package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {

	r := gin.Default()
	r.GET("/toProtobuf", func(context *gin.Context) {
		label := "ss"
		reps := []int64{1, 3}
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		context.ProtoBuf(200, data)
	})
	r.Run(":8090")
}
