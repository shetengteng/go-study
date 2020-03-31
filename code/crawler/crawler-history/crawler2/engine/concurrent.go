package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	// 设置工作的channel
	e.Scheduler.SetMasterWorkerChan(in)

	// 初始化，将种子请求放入
	for _, req := range seeds {
		// 在in 中放入 request
		e.Scheduler.Submit(req)
	}
	// 创建多个goroutine
	// 从in中获取request，将结果返回到out
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	// 主goroutine ，从out中获取数据进行处理
	var count = 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("#%d got item %v ",count, item)
			count ++
		}

		// 如果submit没有单独开辟一个goroutine,那么放入多个request时，主goroutine会阻塞
		// result:=<-out 就不会执行，但是worker中会向out中放入result，则会阻塞，造成死锁
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}
