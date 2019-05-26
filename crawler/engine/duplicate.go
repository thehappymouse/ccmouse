package engine

import "sync"

func init()  {

}
// 去重复 visitedUrls 需要存盘
var visitedUrls = make(map[string]bool)

var mu sync.Mutex
func IsDuplicate(url string) bool {
	mu.Lock()
	defer  mu.Unlock()
	if visitedUrls[url] {
		//log.Error("重重的url:%s", url)
		return true
	}
	visitedUrls[url] = true
	return false
}
