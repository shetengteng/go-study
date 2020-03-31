package engine

type Request struct {
	Url        string                   // 输入的url
	ParserFunc func([]byte) ParseResult // 解析函数，不同业务返回不同的解析结果
}

type ParseResult struct {
	Requests []Request
	Items    []interface{} // city中是city的名称
}

// 定义一个空的ParserFunc
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
