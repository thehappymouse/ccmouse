package engine

import (
	"dali.cc/ccmouse/crawler/model"
	"github.com/gpmgo/gopm/modules/log"
)

// 并发的引擎
// 引擎将请求发送给调度器，调度器纷发给workers, workers的结果再返回给引擎
// 所有worker 引用一个源
type ConcurrentEngine struct {
	MaxWorkerCount int
	Scheduler
}
type Scheduler interface {
	Submit(request Request)
	GetWorkerChan() chan Request

	Run()
	Ready
}
type Ready interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.MaxWorkerCount; i++ {
		e.createWorker(e.Scheduler.GetWorkerChan(), out, e.Scheduler)
	}
	for _, r := range seed {
		if IsDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	// 只打印用户
	var profileCount = 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				profileCount++
				log.Warn("Got Item: #%d %v", profileCount, item)
			}
		}
		for _, r := range result.Requests {
			if IsDuplicate(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
		}
	}

}
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, s Ready) {

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
