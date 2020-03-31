package parser

import (
	"learngo/crawler-history/crawler4history/crawler4/engine"
	"regexp"
)

// 从城市列表的某个城市的网页中解析用户列表，得到各个用户的信息
var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

// 解析城市列表，得到用户的页面
func ParseCity(contents []byte) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1) // [][][]byte
	result := engine.ParseResult{}

	for _, m := range matches {
		userName := string(m[2]) // m会是最后一个人的信息，因此打印始终是这一个人的，需要在这里拷贝出
		url := string(m[1])
		//result.Items = append(result.Items, "User:"+userName) // 存放用户名称 省略，不做处理
		// add
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: url, // 存放各个用户的URL
				ParserFunc: func(contents []byte) engine.ParseResult {
					// 解析人物，将用户的名称添加，使用闭包
					return ParseProfile(contents, url, userName)
				},
			},
		)
	}

	// 解析其他信息，非用户信息
	matches = cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			},
		)
	}

	return result
}
