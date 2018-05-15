package engine

// 解析器接口，支持序列化
type Parser interface {
	Parse(body []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 请求，包括URL和指定的解析函数
type Request struct {
	Url    string
	Parser Parser
}

// 解析结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// 空解析方法
func NilParseFunc([]byte) ParseResult {
	return ParseResult{}
}

// 一个页面对象
type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

// 空解析器
type NilParser struct {
}

func (NilParser) Parse(body []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// 将函数包装成触解析器
type ParserFunc func(body []byte, url string) ParseResult
type FuncParse struct {
	parser ParserFunc
	name   string
}

func (f *FuncParse) Parse(body []byte, url string) ParseResult {
	return f.Parse(body, url)
}

func (f *FuncParse) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParse {
	return &FuncParse{
		parser: p,
		name:   name,
	}
}
