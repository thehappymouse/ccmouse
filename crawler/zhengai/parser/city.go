package parser

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
)

var (
	profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)

	cityUrlRe  = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(contents []byte) engine.ParseResult {
	rs := engine.ParseResult{}

	match := profileRe.FindAllSubmatch(contents, -1)

	for _, m := range match {
		name := string(m[2])
		rs.Items = append(rs.Items, "User " + name)
		rs.Requests = append(rs.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	// 取本页面其它城市链接
	match = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		rs.Requests = append(rs.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return rs
}
