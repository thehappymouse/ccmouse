package engine

import "github.com/gpmgo/gopm/modules/log"

// 并发的引擎
// 引擎将请求发送给调度器，调度器纷发给workers, workers的结果再返回给引擎
// 所有worker 引用一个源
type ConcurrentEngine struct {
	MaxWorkerCount int
	Scheduler
}
type Scheduler interface {
	Submit(request ...Request)
	ConfigMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.MaxWorkerCount; i++ {
		e.createWorker(out, e.Scheduler)
	}

	e.Scheduler.Submit(seed...)
	var count = 0
	for {
		result := <-out
		for _, item := range result.Items {
			count++
			log.Warn("Got Item: #%d %v",count, item)
		}
		e.Scheduler.Submit(result.Requests...)
	}

}
func (e *ConcurrentEngine) createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
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
