package engine

import (
	"learngo/crawler-history/crawler2history/crawler2/fetcher"
	"log"
)

func createWorker(in chan Request, out chan ParseResult) {
	// 开启一个goroutine 进行爬取工作
	go func() {
		for {
			// 获取一个输入，从in中获取
			request := <-in
			result, err := doWork(request)
			if err != nil {
				continue
			}
			// 得到结果输出
			out <- result
		}
	}()
}

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
