package main

import (
	"log"
	"dali.cc/ccmouse/crawler_distributed/rpcsupport"
	"fmt"
	"dali.cc/ccmouse/crawler_distributed/config"
	"dali.cc/ccmouse/crawler_distributed/worker"
	"flag"
)

var port = flag.Int("port", config.WorkerPort0, "请输入工作端口号(默认10086）")

func main() {
	flag.Parse()
	port2 := fmt.Sprintf(":%d", *port)

	fmt.Println("Worker Server Start At:", port2)
	log.Fatal(rpcsupport.ServeRpc(port2,
		worker.CrawlService{}))
}
