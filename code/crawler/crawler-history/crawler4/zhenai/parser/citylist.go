package parser

import (
	"learngo/crawler-history/crawler4history/crawler4/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`

// 解析城市
func ParseCityList(contents []byte) engine.ParseResult {

	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(contents, -1) // [][][]byte

	result := engine.ParseResult{}

	limit := 1 // 测试，获取一个城市

	for _, m := range matches {
		//result.Items = append(result.Items, "City:"+string(m[2])) // 存放城市名称
		// add
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]), // 存放URL
				ParserFunc: ParseCity,
			},
		)

		limit--
		if limit == 0 {
			break
		}

	}
	return result
}
