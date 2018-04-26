package engine

import (
	"dali.cc/ccmouse/crawler/fetcher"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)
// 单任务版引擎
type SimpleEngine struct {

}
func (e SimpleEngine) Run(queue ...Request) {

	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]

		results, err := e.worker(r)
		if err != nil {
			continue
		}

		queue = append(queue, results.Requests...)
		for _, item := range results.Items {
			log.Warn("Got Item:  %v", item)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (e SimpleEngine) worker (r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	log.Warn("Fetching %s", r.Url)
	if err != nil {
		log.ERROR("请求[%s]失败：%s", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}