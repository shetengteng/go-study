package engine

import (
	"learngo/crawler-history/crawler1/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, req := range seeds {
		requests = append(requests, req)
	}

	for len(requests) > 0 {

		// 获取头部request
		req := requests[0]
		// 去除头部request
		requests = requests[1:]

		log.Printf("Fetching url %s ", req.Url)

		// 对每个request进行获取数据
		body, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Printf("Fetcher error fetching url %s : %v", req.Url, err)
			continue
		}
		// 解析每个request的数据
		parseResult := req.ParserFunc(body)

		// 将解析的结果中的request再放入队列中
		requests = append(requests, parseResult.Requests...)

		// 打印每次新增的request
		for _, item := range parseResult.Items {
			log.Printf("got item %v ", item)
		}

	}
}
