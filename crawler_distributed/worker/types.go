package worker

import (
	"dali.cc/ccmouse/crawler/engine"
	"github.com/pkg/errors"
	"dali.cc/ccmouse/crawler/zhengai/parser"
	"github.com/gpmgo/gopm/modules/log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parse.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}
func SerializeResult(r engine.ParseResult) (p ParseResult) {
	p.Items = r.Items
	for _, req := range r.Requests {
		p.Requests = append(p.Requests, SerializeRequest(req))
	}
	return p
}

// todo 整更名称，和 crawler同步
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case "ParseCity":
		return engine.NewFuncParser(parser.ParseCity, p.Name), nil
	case "ParseCityList":
		return engine.NewFuncParser(parser.ParseCityList, p.Name), nil
	case "ProfileParser":
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, errors.New("invalid args for profileParser")
		}

	case "NilParser":
		return engine.NilParse{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
func DeserializeRequest(r Request) (engine.Request, error) {
	parse, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	req := engine.Request{
		Url:   r.Url,
		Parse: parse,
	}
	return req, nil
}
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		ereq, err := DeserializeRequest(req)
		if err != nil {
			log.Warn("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, ereq)
	}
	return result
}

