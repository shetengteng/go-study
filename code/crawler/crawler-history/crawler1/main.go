package main

import (
	"learngo/crawler-history/crawler1history/crawler1/engine"
	"learngo/crawler-history/crawler1history/crawler1/zhenai/parser"
)

func main() {

	seed := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}

	engine.Run(seed)
}
