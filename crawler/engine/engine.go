package engine

import (
	"dali.cc/ccmouse/crawler/fetcher"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

func Run(queue ...Request) {

	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]

		body, err := fetcher.Fetch(r.Url)
		log.Warn("Fetching %s", r.Url)
		if err != nil {
			log.Warn("请求[%s]失败：%s", r.Url, err)
			continue
		}

		results := r.ParseFunc(body)
		queue = append(queue, results.Requests...)
		for _, item := range results.Items {
			log.Warn("Got Item:  %v", item)
		}
		time.Sleep(time.Millisecond * 100)
	}
}
