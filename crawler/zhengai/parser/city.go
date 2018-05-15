package parser

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)

	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}

	match := profileRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),
			Parse: NewProfileParser(string(m[2])),
		})
	}
	// 取本页面其它城市链接
	match = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),
			Parse: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return rs
}
