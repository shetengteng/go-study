package main

import (
	"fmt"
	"net/http"
)

// 实现一个http.Handler接口，用来处理请求
func hello(w http.ResponseWriter, req *http.Request) {
	// 向w中输出hello字符串
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			// 输出header头部的信息
			fmt.Fprintf(w, "%v : %v\n", name, h)
		}
	}
}

func main() {
	// 对get请求进行解析
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
