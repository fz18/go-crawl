package scheduler

import "awesomeProject/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
	return
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workChan <- r }()
}

func (s *SimpleScheduler) ConfigureWorkChan(c chan engine.Request) {
	s.workChan = c
}
