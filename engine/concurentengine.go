package engine

import (
	"awesomeProject/fetcher"
	"log"
)

type Scheduler interface {
	Submit(Request)
	Run()
	WorkerReady(chan Request)
	WorkChan() chan Request
}

type Concurrentengine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}

func (e *Concurrentengine) Run(seek ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()
	for i := 0; i < e.WorkCount; i++ {
		CreateWork(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seek {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	//fmt.Printf("Fetch Url: %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error: %s", err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
