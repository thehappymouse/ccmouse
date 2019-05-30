package engine

import (
	"github.com/rs/zerolog/log"
	"time"
)

func init() {

}

// 去重复 visitedUrls 需要存盘
var urlStore *JsonStore

func IsDuplicate(url string) bool {
	if urlStore.Get(url) == nil {
		urlStore.Set(url, true)
		return false
	}
	return true
}

func InitDuplicateStore() {
	urlStore = CreateJsonStore("urls.json")
	// 需要清理出所有 非 html 结束的页面（可能包含更新）
	go func() {
		for {
			time.Sleep(time.Second * 10)
			gs := urlStore.WriteDisk()
			log.Warn().Msgf("URL数据已存盘, 已访问数量[%d]", gs)
		}
	}()
}
