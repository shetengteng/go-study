package main

import (
	"log"
	"net/http"
	"stt"
)

func main() {

	router := stt.New()

	router.GET("/", func(c *stt.Context) {
		c.HTML(http.StatusOK, "<h1>hello</h1>")
	})

	v1 := router.Group("/v1")
	// 代码块
	{
		v1.GET("/", func(c *stt.Context) {
			c.HTML(http.StatusOK, "<h1>v1</h1>")
		})
		v1.GET("/hello", func(c *stt.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *stt.Context) {
			c.String(http.StatusOK, "hello %s at %s \n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *stt.Context) {
			c.JSON(http.StatusOK, stt.M{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	log.Fatal(router.Run(":9999"))
}
