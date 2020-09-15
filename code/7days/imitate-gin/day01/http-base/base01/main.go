package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// 没有异常时，会阻塞监听，一旦有异常，会返回error信息并打印，退出
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	// 打印路径信息
	// 将输出的信息返回到writer流中
	fmt.Fprintf(writer, "URL.Path= %q\n", request.URL.Path)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = [%q]\n", k, v)
	}
}
