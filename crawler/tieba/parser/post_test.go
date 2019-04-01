package parser

import (
	"dali.cc/ccmouse/crawler/engine"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	body, err := ioutil.ReadFile("post_test_data.html")
	if err != nil {
		panic(err)
	}
	url := "http://tieba.baidu.com/p/4639254365?pn=9"
	result := NewPostParser("冰之泪").Parse(body, url)

	profile := result.Items[0]
	right := engine.Item{
		Url:  url,
		Id:   "http://tieba.baidu.com/p/4639254365?pn=10",
		Type: "tieba",
	}

	if result.Requests[0].Url != right.Id {
		t.Errorf("不相同的会员信息: \n %v : \n %v", profile.Id, right.Id)
	}
	if len(result.Requests) == 0 {
		t.Errorf("猜你喜欢，不应该为空")
	}
}
