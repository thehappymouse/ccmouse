package engine

// 请求，包括URL和指定的解析函数
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}
// 解析结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParseFunc([]byte) ParseResult  {
	return ParseResult{}
}
