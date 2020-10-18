package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/go",myHander)
	// 第二个参数是回调函数
	http.ListenAndServe("127.0.0.1:8090",nil)
}

func myHander(writer http.ResponseWriter, request *http.Request){
	fmt.Println(request.RemoteAddr+" connect succeed")
	fmt.Println("method:",request.Method)
	fmt.Println("url:",request.URL.Path)
	fmt.Println("header:",request.Header)
	fmt.Println("body:",request.Body)
	writer.Write([]byte("hello world"))
}
