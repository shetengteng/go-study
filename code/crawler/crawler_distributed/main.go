package main

import (
	"errors"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
	itemsaver "learngo/crawler_distributed/persist/client"
	"learngo/crawler_distributed/rpcsupport"
	worker "learngo/crawler_distributed/worker/client"
	"log"
	"net/rpc"
)

// 可以设置命令行参数读取
func main() {

	// RPC to save
	saver, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	var hosts = []string{
		config.WorkerPort0,
		config.WorkerPort1,
	}

	pool := createClientPool(hosts)

	// RPC to worker
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         saver,
		RequestProcessor: processor,
	}

	// 种子页面
	seed := engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	}

	e.Run(seed)
}

// 创建连接池
// 如果启动多个worker，需要一个列表进行轮询选择
func createClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("error connecting to %s", h)
			continue

		}
		clients = append(clients, client)
	}

	if len(clients) == 0 {
		panic(errors.New("len of client is 0"))
	}

	clientChan := make(chan *rpc.Client)
	// 通过消息传递给worker的goroutine
	go func() {
		for {
			for _, c := range clients {
				clientChan <- c
			}
		}
	}()
	return clientChan
}
