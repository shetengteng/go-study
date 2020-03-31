package scheduler

import "learngo/crawler-history/crawler2history/crawler2/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) SetMasterWorkerChan(in chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {

	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {

			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 当request 和 worker 都不为空则初始化
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				// 成功将request 放入worker的chan中
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()
}
