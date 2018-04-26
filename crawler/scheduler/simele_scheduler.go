package scheduler

import "dali.cc/ccmouse/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(worker chan engine.Request) {
	s.WorkerChan = worker
}

func (s *SimpleScheduler) Submit(request ...engine.Request) {
	go func() {
		for _, r := range request {
			s.WorkerChan <- r
		}
	}()
}
