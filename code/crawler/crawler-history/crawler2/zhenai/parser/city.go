package parser

import (
	"learngo/crawler-history/crawler2history/crawler2/engine"
	"regexp"
)

// 从城市列表的某个城市的网页中解析用户列表，得到各个用户的信息
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

// 解析城市列表，得到用户的页面
func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRe)
	matches := compile.FindAllSubmatch(contents, -1) // [][][]byte

	result := engine.ParseResult{}

	for _, m := range matches {

		userName := string(m[2])

		result.Items = append(result.Items, "User:"+userName) // 存放用户名称

		// add
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]), // 存放各个用户的URL
				ParserFunc: func(contents []byte) engine.ParseResult {
					// 解析人物，将用户的名称添加，使用闭包
					return ParseProfile(contents, userName)
				},
			},
		)
	}
	return result
}
