package scheduler

import "awesomeProject/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueueScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request

			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]

			}
		}
	}()
}
