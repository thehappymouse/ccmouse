package engine

import (
	"dali.cc/ccmouse/crawler/fetcher"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

func Run(seeds ...Request) {
	q := []Request{}
	for _, r := range seeds {
		q = append(q, r)
	}
	for len(q) > 0 {
		r := q[0]
		q = q[1:]

		body, err := fetcher.Fetch(r.Url)
		log.Warn("Fetching %s", r.Url)
		if err != nil {
			log.Warn("请求[%s]失败：%s", r.Url, err)
			continue
		}

		ps := r.ParseFunc(body)
		q = append(q, ps.Requests...)
		for _, item := range ps.Items {
			log.Warn("Got Item %v", item)
		}
		time.Sleep(time.Millisecond * 10)
	}
}
