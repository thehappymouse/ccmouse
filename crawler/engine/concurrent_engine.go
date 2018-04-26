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
}

func (e *ConcurrentEngine) Run(seed ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	for i := 0; i < e.MaxWorkerCount; i++ {
		e.createWorker(in, out)
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
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
