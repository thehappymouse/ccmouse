package main

import (
	"dali.cc/ccmouse/crawler/engine"
	parser2 "dali.cc/crawler_tieba/parser"
	"dali.cc/ccmouse/crawler/zhengai/parser"
)


func tieBa() {

	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:   "http://tieba.baidu.com/p/5571119322",
			Parse: parser2.NewPostParser("first"),
		},
		{
			Url:   "http://tieba.baidu.com/p/5002958965",
			Parse: parser2.NewPostParser("two"),
		},
		{
			Url:   "https://tieba.baidu.com/p/4575246659",
			Parse: parser2.NewPostParser("three"),
		},
	}
	e := engine.SimpleEngine{}	//单机
	e.Run(seed...)
}

func zhenai()  {
	//itemChan, err := persist.ItemSaver("profiles")
	//if err != nil {
	//	panic(err)
	//}
	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:   "http://www.zhenai.com/zhenghun/beijing",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
		//{
		//	Url:       "http://www.zhenai.com/zhenghun/henan",
		//	ParseFunc: parser.ParseCity,
		//},
		//{
		//	Url:   "http://www.zhenai.com/zhenghun",
		//	Parse: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
		//},
	}
	e := engine.SimpleEngine{}	//单机
	//e := engine.ConcurrentEngine{
	//	MaxWorkerCount:   100,
	//	Scheduler:        &scheduler.QueuedScheduler{},
	//	ItemChan:         itemChan,
	//	RequestWorker: engine.Worker,
	//}
	e.Run(seed...)
}

// 单机，并发
func main() {
	tieBa()
}
