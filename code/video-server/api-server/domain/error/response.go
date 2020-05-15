package error

// 错误消息返回体
type Response struct {
	HttpSC int
	Error  Info
}

// 错误消息
type Info struct {
	Message string `json:"error"`
	Code    string `json:"error_code"`
}
