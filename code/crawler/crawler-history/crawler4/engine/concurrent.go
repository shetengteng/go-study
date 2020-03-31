package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)

	// 设置工作的channel
	e.Scheduler.Run()

	// 初始化，将种子请求放入
	for _, req := range seeds {
		// 在in 中放入 request
		e.Scheduler.Submit(req)
	}
	// 创建多个goroutine
	// 从in中获取request，将结果返回到out
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 主goroutine ，从out中获取数据进行处理
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("#%d got item %v ", count, item)
			// 开辟一个goroutine进行保存操作
			go func() { e.ItemChan <- item }()
		}

		// 如果submit没有单独开辟一个goroutine,那么放入多个request时，主goroutine会阻塞
		// result:=<-out 就不会执行，但是worker中会向out中放入result，则会阻塞，造成死锁
		for _, request := range result.Requests {
			// 去重
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

// 用于去重
var dict = make(map[string]bool)

func isDuplicate(url string) bool {
	_, ok := dict[url]
	if ok {
		return true
	}
	dict[url] = true
	return false
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	// 开启一个goroutine 进行爬取工作
	// 每一个worker 一个worker request channel
	go func() {
		for {
			// 说明该worker准备好了
			ready.WorkerReady(in)
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
