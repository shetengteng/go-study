package engine

type Scheduler interface {
	Submit(Request)
	SetMasterWorkerChan(chan Request)
}