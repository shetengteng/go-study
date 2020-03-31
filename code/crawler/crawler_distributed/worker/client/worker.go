package client

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/worker"
	"net/rpc"
)

// 定义一个worker的操作数据的客户端

// 将clients从外部传递，通过chan，避免从clients[] 中上锁获取资源
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	// 返回处理函数
	return func(req engine.Request) (engine.ParseResult, error) {

		// 序列化request
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult

		// 每次从clientChan中获取数据
		c := <-clientChan

		// 调用rpc
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		// 反序列化结果
		return worker.DeserializeResult(sResult), nil
	}
}
