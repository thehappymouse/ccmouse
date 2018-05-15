package main

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/scheduler"
	"dali.cc/ccmouse/crawler/persist"
	"dali.cc/ccmouse/crawler/zhengai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("profiles")
	if err != nil {
		panic(err)
	}
	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:    "http://www.zhenai.com/zhenghun/beijing",
			Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
	}
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		MaxWorkerCount: 200,
		Scheduler:      &scheduler.QueuedScheduler{},
		ItemChan:       itemChan,
	}
	e.Run(seed...)
}
