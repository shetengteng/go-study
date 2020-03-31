package main

import (
	"learngo/crawler-history/crawler2history/crawler2/engine"
	"learngo/crawler-history/crawler2history/crawler2/scheduler"
	"learngo/crawler-history/crawler2history/crawler2/zhenai/parser"
)

func main() {

	// 种子页面
	seed := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{}, // 设置为简单调度，只有1个进1个出
		WorkerCount: 10,                           // 同时工作的协程数量是10
	}

	e.Run(seed)
}
