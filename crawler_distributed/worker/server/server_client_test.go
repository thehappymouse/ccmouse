package main

import (
	"testing"
	"dali.cc/ccmouse/crawler_distributed/rpcsupport"
	"dali.cc/ccmouse/crawler_distributed/worker"
	"time"
	"dali.cc/ccmouse/crawler_distributed/config"
	"fmt"
)

func TestCrawService(t *testing.T) {
	host := ":9003"

	go func() {
		rpcsupport.ServeRpc(host, worker.CrawlService{})
	}()
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1077868794",
		Parser: worker.SerializedParser{
			Name: "ProfileParser",
			Args: "冰之泪",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
