package main

import (
	"learngo/crawler-history/crawler4history/crawler4/engine"
	"learngo/crawler-history/crawler4history/crawler4/persist"
	"learngo/crawler-history/crawler4history/crawler4/scheduler"
	"learngo/crawler-history/crawler4history/crawler4/zhenai/parser"
)

func main() {

	saver, err := persist.ItemSaver("dating_profile")

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    saver,
	}

	// 种子页面
	seed := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}

	e.Run(seed)
}
