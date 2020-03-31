package main

import (
	"learngo/crawler-history/crawler3history/crawler3/engine"
	"learngo/crawler-history/crawler3history/crawler3/scheduler"
	"learngo/crawler-history/crawler3history/crawler3/zhenai/parser"
)

func main() {

	// 种子页面
	seed := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}

	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.QueuedScheduler{},
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,                           // 同时工作的协程数量是10
	}

	e.Run(seed)
}
