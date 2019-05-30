package engine

// 去重复 visitedUrls 需要存盘
var urlStore = &JsonStore{}

func IsDuplicate(url string) bool {
	if urlStore.Get(url) == nil {
		urlStore.Set(url, true)
		return false
	}
	return true
}

// 设置重复的store
func SetDuplicateStore(s *JsonStore) {
	urlStore = s
}
