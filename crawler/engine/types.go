package engine

type ParserFunc func(body []byte, url string) ParseResult

// 请求，包括URL和指定的解析函数
type Request struct {
	Url       string
	ParseFunc ParserFunc
}
// 解析结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}
// 空解析方法
func NilParseFunc([]byte) ParseResult  {
	return ParseResult{}
}

// 一个页面对象
type Item struct {
	Url string
	Id string
	Type string
	Payload interface{}
}