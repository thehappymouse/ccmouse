package main

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/zhengai/parser"
	"dali.cc/ccmouse/crawler/scheduler"
	"dali.cc/ccmouse/crawler/persist"
)

var (
	startUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	//seed := engine.Request{
	//	Url:       startUrl,
	//	ParseFunc: parser.ParseCityList,
	//}
	beijingSeed := engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/beijing",
		ParseFunc: parser.ParseCity,
	}
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		MaxWorkerCount: 10,
		Scheduler:      &scheduler.QueuedScheduler{},
		ItemChan:       persist.ItemSaver(),
	}
	e.Run(beijingSeed)
}
