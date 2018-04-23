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
	results := ParseCityList(body)

	var resultSize = 470

	if l := len(results.Items); l != resultSize {
		t.Errorf("计算结果是: %d, 应该是: %d", l, resultSize)
	}

	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectedCities = []string{
		"阿坝","阿克苏","阿拉善盟",
	}

	for i, url := range expectedUrls {
		if l := results.Requests[i].Url; l != url {
			t.Errorf("肯定包含的url # %d: %s, 但计算得到的是: %s", i, url, l)
		}
	}
	for i, city := range expectedCities {
		if l := results.Items[i]; l != city {
			t.Errorf("肯定包含的城市 # %d: %s, 但计算得到的是: %s", i, city, l)
		}
	}
}
