package stt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type M map[string]interface{}

type Context struct {
	// 原生对象
	Writer http.ResponseWriter
	// 注意：这里是请求对象的引用
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int

	// 存储url中匹配的参数
	Params map[string]string
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 获取form表单中的数据
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

// 获取url中的数据
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// 在http头部写入状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 在http头部设置key-value
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 在输出流中输出byte数组
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 在输出流中输出string
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	// 此处传递values...表示将数组打散，一个一个传递，否则传递的是数组的对象，而非数组的元素
	c.Data(code, []byte(fmt.Sprintf(format, values...)))
}

// 在输出流中输出html页面
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Data(code, []byte(html))
}

// 在输出流中输出JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Context-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
