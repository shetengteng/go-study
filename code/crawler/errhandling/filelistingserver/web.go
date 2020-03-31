package main

import (
	"learngo/errhandling/filelistingserver/filelisting"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

func main() {

	// 通过url访问文件
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

// 给HandlerFileList方法设置一个类型，使用函数式编程，函数做为参数和返回值
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 处理异常
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// 增加recover的异常处理，处理panic
		defer func() {
			r := recover()
			if r != nil {
				log.Warn("panic :%v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request: %s", err.Error())

			// 如果错误是userError，处理自定义异常
			if userErr, ok := err.(filelisting.UserError); ok {
				http.Error(writer,
					userErr.Message(), // 获取Message消息
					http.StatusBadRequest)
				return
			}

			// 处理系统异常
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
