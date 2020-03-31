package engine

import (
	"learngo/crawler-history/crawler4history/crawler4/fetcher"
	"log"
)

// 对页面进行访问，并返回解析结果
func doWork(req Request) (ParseResult, error) {

	log.Printf("Fetching url %s ", req.Url)

	// 对每个request进行获取数据
	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url %s : %v", req.Url, err)
		return ParseResult{}, err
	}
	// 解析每个request的数据
	return req.ParserFunc(body), nil
}
