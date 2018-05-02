package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"log"
	"dali.cc/ccmouse/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	ch := make(chan engine.Item, 1024)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for item := range ch {
			itemCount++
			log.Printf("Item Saver: Got Item #%d: %v", itemCount, item)
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: Save error: %s", err)
			}
		}
	}()
	return ch, nil
}

// 返回存储的ID
func Save(client *elastic.Client, index string, item engine.Item) (error) {

	if  item.Type == "" {
		return errors.New("item.type 不能为空")
	}
	_, err := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).Do(context.Background())
	return err
}
