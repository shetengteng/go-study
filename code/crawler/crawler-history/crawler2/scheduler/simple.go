package scheduler

import "learngo/crawler-history/crawler2history/crawler2/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) SetMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}
func (s *SimpleScheduler) Submit(request engine.Request) {
	// 每个request创建一个goroutine，用于存放request，否则会有死锁
	go func() { s.workerChan <- request }()
}
