package engine

type ParserFunc func(body []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 请求，包括URL和指定的解析函数
type Request struct {
	Url   string
	Parse Parser
}

// 解析结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// 空解析方法
func NilParseFunc(body []byte, url string) ParseResult {
	return ParseResult{}
}

// 一个页面对象
type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

type NilParse struct {
}

func (NilParse) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParse) Serialize() (name string, args interface{}) {
	return "NilParse", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (p *FuncParser) Parse(contents []byte, url string) ParseResult {
	return p.parser(contents, url)
}

func (p *FuncParser) Serialize() (name string, args interface{}) {
	return p.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
