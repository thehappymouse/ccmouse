package main

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/scheduler"
	"dali.cc/ccmouse/crawler/zhengai/parser"
	itemsaver "dali.cc/ccmouse/crawler_distributed/persist/client"
	"dali.cc/ccmouse/crawler_distributed/config"
	worker "dali.cc/ccmouse/crawler_distributed/worker/client"
)

func main() {

	itemChan, err := itemsaver.ItemSaver(config.ItemSaverPort)
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:   "http://www.zhenai.com/zhenghun/beijing",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
	}

	e := engine.ConcurrentEngine{
		MaxWorkerCount: 200,
		Scheduler:      &scheduler.QueuedScheduler{},
		ItemChan:       itemChan,
		//RequestWorker:  engine.Worker,
		RequestWorker: processor,
	}
	e.Run(seed...)
}
