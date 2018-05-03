package client

import (
	"log"
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	ch := make(chan engine.Item, 1024)

	rpc, err := rpcsupport.NewClient(host)

	go func() {
		itemCount := 0
		for item := range ch {
			itemCount++
			log.Printf("Item Saver: Got Item #%d: %v", itemCount, item)

			result := ""
			rpc.Call("ItemSaverService.Save", item, &result)

			if err != nil {
				log.Printf("Item Saver: Save error: %s", err)
			}
		}
	}()
	return ch, nil
}
