package main

import (
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/scheduler"
	"dali.cc/ccmouse/crawler/zhengai/parser"
	itemsaver "dali.cc/ccmouse/crawler_distributed/persist/client"
	"dali.cc/ccmouse/crawler_distributed/config"
	worker "dali.cc/ccmouse/crawler_distributed/worker/client"
	"net/rpc"
	"dali.cc/ccmouse/crawler_distributed/rpcsupport"
	"github.com/gpmgo/gopm/modules/log"
	"flag"
	"strings"
)

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Warn("error connection to %s : %s", h, err)
		} else {
			clients = append(clients, client)
			log.Warn("connected  to %s", h)
		}
	}
	out := make(chan *rpc.Client)
	// 持续纷发客户端
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}

var hosts = flag.String("hosts", "", "多个工作节点的端口，以逗号分隔,例如 :9002,:9003")

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(config.ItemSaverPort)
	if err != nil {
		panic(err)
	}
	log.Warn("connected item saver at : %v", config.ItemSaverPort)

	pool := createClientPool(strings.Split(*hosts, ","))

	processor := worker.CreateProcessor(pool)

	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:       "http://www.zhenai.com/zhenghun/henan",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
		{
			Url:   "http://www.zhenai.com/zhenghun/beijing",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
	}

	e := engine.ConcurrentEngine{
		MaxWorkerCount: 100,
		Scheduler:      &scheduler.QueuedScheduler{},
		ItemChan:       itemChan,
		//RequestWorker:  engine.Worker,	//单work
		RequestWorker: processor, //rpc work
	}
	e.Run(seed...)
}
