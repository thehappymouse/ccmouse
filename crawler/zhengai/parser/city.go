package parser

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
)

const cityRegex = `<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	rs := engine.ParseResult{}

	reg := regexp.MustCompile(cityRegex)
	match := reg.FindAllSubmatch(contents, -1)

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

	return rs
}
