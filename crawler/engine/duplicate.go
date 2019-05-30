package engine

type DuplicateChecker interface {
	IsDuplicate(url string) bool
}

type FileDuplicateChecker struct {
	urlStore *JsonStore
}

func (t *FileDuplicateChecker) IsDuplicate(url string) bool {
	if t.urlStore.Get(url) == nil {
		t.urlStore.Set(url, true)
		return false
	}
	return true
}

var urlChecker DuplicateChecker

func IsDuplicate(url string) bool {

	return urlChecker.IsDuplicate(url)
}

func SetDuplicateChecker(c DuplicateChecker)  {
	urlChecker = c
}

// 设置重复的store
func SetDuplicateStore(s *JsonStore) {
	urlChecker = &FileDuplicateChecker{
		urlStore: s,
	}
}
