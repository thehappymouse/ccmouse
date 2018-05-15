package parser

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
)

const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}

	reg := regexp.MustCompile(cityListRegex)
	match := reg.FindAllSubmatch(contents, -1)
	for _, m := range match {
		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),
			Parse: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return rs
}
