package main

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/zhengai/parser"
	"dali.cc/ccmouse/crawler/scheduler"
)

var (
	startUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	seed := engine.Request{
		Url:       startUrl,
		ParseFunc: parser.ParseCityList,
	}
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		MaxWorkerCount: 200,
		Scheduler: &scheduler.SimpleScheduler{},
		//Scheduler: &scheduler.QueuedScheduler{},
	}
	e.Run(seed)
}
