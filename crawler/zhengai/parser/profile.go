package parser

// 解析会员信息

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
	"dali.cc/ccmouse/crawler/model"
	"reflect"
	"github.com/gpmgo/gopm/modules/log"
)

var regexs = map[string]*regexp.Regexp{
	"Age":        regexp.MustCompile(`<td><span class="label">年龄：</span>(.+)</td>`),
	"Sex":        regexp.MustCompile(`<td><span class="label">性别：</span><span field="">(.+)</span></td>`),
	"Height":     regexp.MustCompile(`<td><span class="label">身高：</span>(.+)</td>`),
	"Marriage":   regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`),
	"Edu":        regexp.MustCompile(`<td><span class="label">学历：</span>(.+)</td>`),
	"Job":        regexp.MustCompile(`<td><span class="label">职业：.*</span>([^<]+)</td>`),
	"JobAddress": regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`),
	"Child":      regexp.MustCompile(`<td><span class="label">有.+孩子：</span>(.+)</td>`),
	"Income":     regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`),
}
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

// 生成用户解析函数的函数
func ProfileParser(name string) engine.ParserFunc   {
	return func(body []byte, url string) engine.ParseResult {
		return ParseProfile(body, url, name)
	}
}

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
			log.Warn("未能解析的属性：%s", k)
		}
	}
	item := engine.Item{
		Url:     url,
		Payload: profile,
		Type:    "zhenai",
		Id: extractString([]byte(url), idUrlRe),
	}
	rs.Items = []engine.Item{item}
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
