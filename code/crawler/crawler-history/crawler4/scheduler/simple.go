package scheduler

import "learngo/crawler-history/crawler4history/crawler4/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request // 所有的worker共用一个channel
}


func (s *SimpleScheduler) WorkerChan() chan engine.Request{
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.workerChan <- r}()
}

func (s *SimpleScheduler) SetMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}


