package parser

import (
	"testing"
	"io/ioutil"
	"dali.cc/ccmouse/crawler/model"
)

func TestParseProfile(t *testing.T) {
	body, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	results := ParseProfile(body, "冰之泪")
	profile := results.Items[0].(model.Profile)
	right := model.Profile{
		Name: "冰之泪",
		Age:"47岁",
		Height: "170CM",
		Marriage: "离异",
		Edu:"大专",
		Job:"其他职业",
		Child: "有，我们住在一起",
		JobAddress:"陕西宝鸡",
	}
	if profile != right {
		t.Errorf("不相同的会员信息: \n %v : \n %v", profile, right)
	}
}
