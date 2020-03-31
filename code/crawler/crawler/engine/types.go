package engine

type Request struct {
	Url    string // 输入的url
	Parser Parser // Parser是一个接口
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

//=---------------------------------------

// 解析函数，不同业务返回不同的解析结果 ,如何在网络上传输，需要进行序列化
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 定义NilParser
type NilParser struct{}

// 实现接口
func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// 定义一个ParseFunc类型
type ParseFunc func(contents []byte, url string) ParseResult

// 定义一个FuncParser
type FuncParser struct {
	parser ParseFunc
	name   string
}

// 使用工厂函数的方式构建，注意需要返回指针，由于FuncParser的接口实现的方法是指针类型
func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

// 实现接口
func (p *FuncParser) Parse(content []byte, url string) ParseResult {
	return p.parser(content, url)
}

func (p *FuncParser) Serialize() (name string, args interface{}) {
	return p.name, nil
}
