package parser

// 解析会员信息

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
	"dali.cc/ccmouse/crawler/model"
	"reflect"
)

var regexs = map[string]*regexp.Regexp{
	"Age":        regexp.MustCompile(`<td><span class="label">年龄：</span>(.+)</td>`),
	"Sex":        regexp.MustCompile(`<td><span class="label">性别：</span><span field="">(.+)</span></td>`),
	"Height":     regexp.MustCompile(`<td><span class="label">身高：</span>(.+)</td>`),
	"Marriage":   regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`),
	"Edu":        regexp.MustCompile(`<td><span class="label">学历：</span>(.+)</td>`),
	"Job":        regexp.MustCompile(`<td><span class="label">职业：.*</span>([^<]+)</td>`),
	"JobAddress": regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`),
	"HasChild":   regexp.MustCompile(`<td><span class="label">有无孩子：</span>(.+)</td>`),
	"Income":     regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`),
}
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url, name string) engine.ParseResult {
	rs := engine.ParseResult{}

	profile := model.Profile{Name: name}

	v := reflect.ValueOf(&profile).Elem()

	for k, r := range regexs {
		s := extractString(contents, r)
		if s != "" {
			a := v.FieldByName(k)
			if a.IsValid() {
				a.Set(reflect.ValueOf(s))
			}
		} else {
			//log.Warn("未能解析的属性：%s", k)
		}
	}
	item := engine.Item{
		Url:     url,
		Payload: profile,
		Type:    "zhenai",
		Id:      extractString([]byte(url), idUrlRe),
	}
	rs.Items = []engine.Item{item}

	// 取本页面内，猜你喜欢的的
	var guessRe = regexp.MustCompile(`href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)
	ms := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range ms {
		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),
			Parse:  NewProfileParser(string(m[2])),
		})
	}

	// 取本页面其它城市链接
	return rs
}

func extractString(c []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(c)
	if match != nil && len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}


// 生成用户解析函数的函数
//func ProfileParser(name string) engine.ParserFunc   {
//	return func(body []byte, url string) engine.ParseResult {
//		return ParseProfile(body, url, name)
//	}
//}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
