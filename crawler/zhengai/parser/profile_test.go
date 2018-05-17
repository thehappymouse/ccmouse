package parser

import (
	"testing"
	"io/ioutil"
	"dali.cc/ccmouse/crawler/model"
	"dali.cc/ccmouse/crawler/engine"
)

func TestParseProfile(t *testing.T) {
	body, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	url := "http://album.zhenai.com/u/1077868794"
	result := NewProfileParser("冰之泪").Parse(body, url)

	profile := result.Items[0]
	right := engine.Item{
		Url:  url,
		Id:   "1077868794",
		Type: "zhenai",
		Payload: model.Profile{
			Name:       "冰之泪",
			Age:        "47岁",
			Height:     "170CM",
			Marriage:   "离异",
			Income:     "8001-12000元",
			Edu:        "大专",
			Job:        "其他职业",
			Sex:        "男",
			HasChild:   "有，我们住在一起",
			JobAddress: "陕西宝鸡",
		},
	}
	if profile != right {
		t.Errorf("不相同的会员信息: \n %v : \n %v", profile, right)
	}
	if len(result.Requests) == 0 {
		t.Errorf("猜你喜欢，不应该为空")
	}
}
