package parser

// 解析会员信息

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
	"dali.cc/ccmouse/crawler/model"
	"reflect"
	"github.com/gpmgo/gopm/modules/log"
)

//ID, Name, Height, Weight, Job, JobAddress, Edu, Child, Jiguan, Age string

var regexs = map[string]*regexp.Regexp{
	"Age":        regexp.MustCompile(`<td><span class="label">年龄：</span>(.+)</td>`),
	"Height":     regexp.MustCompile(`<td><span class="label">身高：</span>(.+)</td>`),
	"Marriage":   regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`),
	"Edu":        regexp.MustCompile(`<td><span class="label">学历：</span>(.+)</td>`),
	"Job":        regexp.MustCompile(`<td><span class="label">职业：.*</span>([^<]+)</td>`),
	"JobAddress": regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`),
	"Child":      regexp.MustCompile(`<td><span class="label">有.+孩子：</span>(.+)</td>`),
}

func ParseProfile(contents []byte, name string) engine.ParseResult {
	rs := engine.ParseResult{}
	profile := model.Profile{Name:name}

	v := reflect.ValueOf(&profile).Elem()

	for k, r := range regexs {
		match := r.FindSubmatch(contents)

		if match != nil {
			s := string(match[1])

			a := v.FieldByName(k)
			if a.IsValid() {
				a.Set(reflect.ValueOf(s))
			}
		} else {
			log.Warn("未能解析的属性：%s", k)
		}
	}

	rs.Items = []interface{}{
		profile,
	}
	return rs
}
