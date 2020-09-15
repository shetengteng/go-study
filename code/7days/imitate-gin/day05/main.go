package main

import (
	"log"
	"net/http"
	"stt"
	"time"
)

func logger() stt.HandlerFunc {
	return func(c *stt.Context) {
		// start time
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v ", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func middlewareForV2() stt.HandlerFunc {
	return func(c *stt.Context) {
		// start time
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {

	router := stt.New()
	{
		router.Use(logger()) // 全局 middleware
		router.GET("/", func(c *stt.Context) {
			c.HTML(http.StatusOK, "<h1>hello</h1>")
		})
	}

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
		v2.Use(middlewareForV2())
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
