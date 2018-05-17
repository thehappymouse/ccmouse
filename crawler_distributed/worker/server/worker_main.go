package main

import (
	"log"
	"dali.cc/ccmouse/crawler_distributed/rpcsupport"
	"fmt"
	"dali.cc/ccmouse/crawler_distributed/config"
	"dali.cc/ccmouse/crawler_distributed/worker"
)

func main() {
	port := fmt.Sprintf(":%d", config.WorkerPort0)
	log.Fatal(rpcsupport.ServeRpc(port,
		worker.CrawlService{}))
}
