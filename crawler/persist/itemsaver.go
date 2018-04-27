package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"log"
	"dali.cc/ccmouse/crawler/engine"
)

func ItemSaver() chan engine.Item {
	ch := make(chan engine.Item)
	go func() {
		itemCount := 0
		for item := range ch {
			itemCount++
			log.Printf("Item Saver: Got Item #%d: %v", itemCount, item)
			err := save(item)
			if err != nil {
				log.Printf("Item Saver: save error: %s", err)
			}
		}
	}()
	return ch
}

// 返回存储的ID
func save(item engine.Item) (error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return  err
	}
	_, err = client.Index().
		Index("profiles").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).Do(context.Background())
	if err != nil {
		return  err
	}
	return  nil
}
