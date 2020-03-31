package engine

import (
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
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

		parseResult, err := Worker(req)
		if err != nil {
			continue
		}

		// 将解析的结果中的request再放入队列中
		requests = append(requests, parseResult.Requests...)

		// 打印每次新增的request
		for _, item := range parseResult.Items {
			log.Printf("got item %v ", item)
		}
	}
}
