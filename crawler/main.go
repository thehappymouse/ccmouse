package main

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/zhengai/parser"
)

var (
	startUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:       startUrl,
		ParseFunc: parser.ParseCityList,
	})
}
