package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

// 对页面进行访问，并返回解析结果
func Worker(req Request) (ParseResult, error) {

	log.Printf("Fetching url %s ", req.Url)

	// 对每个request进行获取数据
	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url %s : %v", req.Url, err)
		return ParseResult{}, err
	}
	// 解析每个request的数据
	return req.Parser.Parse(body, req.Url), nil
}
