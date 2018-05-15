package main

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/scheduler"
	"dali.cc/ccmouse/crawler/zhengai/parser"
	"dali.cc/ccmouse/crawler_distributed/persist/client"
	"dali.cc/ccmouse/crawler_distributed/config"
)

func main() {

	itemChan, err := client.ItemSaver(config.ItemSaverPort)
	if err != nil {
		panic(err)
	}

	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:       "http://www.zhenai.com/zhenghun/beijing",
			ParseFunc: parser.ParseCity,
		},
		{
			Url:       "http://www.zhenai.com/zhenghun/henan",
			ParseFunc: parser.ParseCity,
		},
		{
			Url:       "http://www.zhenai.com/zhenghun",
			ParseFunc: parser.ParseCityList,
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
