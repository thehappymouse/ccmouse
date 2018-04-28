package parser

import (
	"testing"
	"io/ioutil"
)

// 测试最好不要有依赖
func TestParseCityList(t *testing.T) {
	body, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	results := ParseCityList(body, "")

	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls {
		if l := results.Requests[i].Url; l != url {
			t.Errorf("肯定包含的url # %d: %s, 但计算得到的是: %s", i, url, l)
		}
	}
}
