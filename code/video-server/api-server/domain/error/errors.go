package error

// 错误消息集合
var (
	ErrorRequestBodyParseFailed = Response{400, Info{
		Message: "request body is not correct",
		Code:    "001",
	}}
	ErrorNotAuthUser = Response{401, Info{
		Message: "user authentication failed",
		Code:    "002",
	}}
	ErrorDBOps = Response{500, Info{
		Message: "database operation error",
		Code:    "003",
	}}
	ErrorInteralFailed = Response{500, Info{
		Message: "error of interal failed",
		Code:    "004",
	}}
	ErrorInvalidSession = Response{403, Info{
		Message: "error of invalid session",
		Code:    "005",
	}}
)
