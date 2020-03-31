package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
)

func main() {

	saver, err := persist.ItemSaver("dating_profile")

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         saver,
		RequestProcessor: engine.Worker, // 单机版的调用worker
	}

	// 种子页面
	seed := engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	}

	e.Run(seed)
}
