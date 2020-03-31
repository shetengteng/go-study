package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

// 封装一个userError类型
type UserError string

// 该方法与error接口一致
func (e UserError) Error() string {
	return e.Message()
}

func (e UserError) Message() string {
	return string(e)
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		// 返回自定义错误
		return UserError("path must start with" + prefix)
	}

	path := request.URL.Path[len(prefix):] //去除list前缀，获取文件路径
	file, error := os.Open(path)
	if error != nil {
		//panic(error)
		// 将服务器给外部显示，不友好
		//http.Error(writer,error.Error(),http.StatusInternalServerError)
		// 将异常抛出
		return error
	}
	defer file.Close()
	all, error := ioutil.ReadAll(file)
	if error != nil {
		//panic(error)
		return error
	}
	writer.Write(all)
	return nil
}
