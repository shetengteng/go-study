package main

import (
	"learngo/crawler-history/crawler4history/crawler4/frontend/controller"
	"net/http"
)

func main() {

	// 针对css js文件进行处理，文件服务
	// fileServer，如果有index.html则会直接展示 / 访问 对应的是index.html
	http.Handle("/", http.FileServer(http.Dir("crawler4/frontend/view")))

	// 对处理的请求分配handler
	http.Handle(
		"/search",
		controller.CreateSearchResultHandler("crawler4/frontend/view/template.html"),
	)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
