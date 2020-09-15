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

	router.GET("/hello", func(c *stt.Context) {
		c.String(http.StatusOK, "hello %s , at %s \n", c.Query("name"), c.Path)
	})

	router.POST("/login", func(c *stt.Context) {
		c.JSON(http.StatusOK, stt.M{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	router.GET("/hello/:name", func(c *stt.Context) {
		c.String(http.StatusOK, "hello %s at %s \n", c.Param("name"), c.Path)
	})

	router.GET("/assets/*filepath", func(c *stt.Context) {
		c.JSON(http.StatusOK, stt.M{"filepath": c.Param("filepath")})
	})

	log.Fatal(router.Run(":9999"))
}
