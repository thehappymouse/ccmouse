package engine

import (
	"github.com/gpmgo/gopm/modules/log"
)
// 单任务版引擎
type SimpleEngine struct {

}
func (e *SimpleEngine) Run(queue ...Request) {

	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]

		results, err := worker(r)
		if err != nil {
			continue
		}

		queue = append(queue, results.Requests...)
		for _, item := range results.Items {
			log.Warn("Got Item:  %v", item)
		}
	}
}

