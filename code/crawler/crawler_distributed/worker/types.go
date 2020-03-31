package worker

import (
	"errors"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
	"log"
)

// 定义序列化结构体
// {"ParseCityList",nil},{"ParseProfile",userName}
type SerializedParser struct {
	Name string // FunctionName
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser // 可以在网络中传递，需要将该Request与engine的Request进行转换
}

type ParseResult struct { // 该ParseResult 和 engine的ParseResult进行转换
	Items    []engine.Item
	Requests []Request
}

// 将engine.Request 转换为 Request
// 序列化
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

// 反序列化
func DeserializeRequest(r Request) (engine.Request, error) {

	parser, error := deserializeParser(r.Parser)
	if error != nil {
		return engine.Request{}, error
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	//方式1： 将parser的名字注册到map中，然后通过map查找parser
	//方式2：
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ProfileParser:
		username, ok := p.Args.(string)
		if ok {
			return parser.NewProfileParser(username), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name: " + p.Name)
	}
}
