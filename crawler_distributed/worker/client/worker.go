package client

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler_distributed/rpcsupport"
	"fmt"
	"dali.cc/ccmouse/crawler_distributed/config"
	"dali.cc/ccmouse/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error)  {
	client,err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(request engine.Request) (engine.ParseResult, error) {
		var sReq = worker.SerializeRequest(request)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
