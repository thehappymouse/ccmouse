package parser

// 解析会员信息

import (
	"dali.cc/ccmouse/crawler/engine"
	"regexp"
	"dali.cc/ccmouse/crawler/model"
	"strconv"
)

var ageRegex = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d\d)岁</td>`)

func ParseProfile(contents []byte) engine.ParseResult {

	rs := engine.ParseResult{
	}

	match := ageRegex.FindSubmatch(contents)
	profile := model.Profile{}
	if len(match) > 1 {
		age, _ := strconv.Atoi(string(match[1]))
		profile.Age = age
	}
	rs.Items = []interface{}{
		profile,
	}

	return rs
}
