package engine

type Request struct {
	Url        string                   // 输入的url
	ParserFunc func([]byte) ParseResult // 解析函数，不同业务返回不同的解析结果 ,如何在网络上传输，需要进行序列化
}

type ParseResult struct {
	Requests []Request
	Items    []Item // city中是city的名称
}

type Item struct {
	Url     string
	Id      string
	Type    string // 存储在ES中的type，表示哪个表
	Payload interface{}
}

// 定义一个空的ParserFunc
func NilParser([]byte) ParseResult {
	return ParseResult{}
}

