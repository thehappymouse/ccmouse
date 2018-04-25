package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCity(t *testing.T) {
	body, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	results := ParseCity(body)
	var resultSize = 20
	if l := len(results.Items); l != resultSize {
		t.Errorf("计算结果是: %d, 应该是: %d", l, resultSize)
	}

	var expectedUrls = []string{
		"http://album.zhenai.com/u/1861056016",
		"http://album.zhenai.com/u/106344107",
		"http://album.zhenai.com/u/1077868794",
	}
	var expectedCities = []string{
		"User 因为有你",
		"User 那只是过去",
		"User 一生的回忆",
	}

	for i, url := range expectedUrls {
		if l := results.Requests[i].Url; l != url {
			t.Errorf("肯定包含的url # %d: %s, 但计算得到的是: %s", i, url, l)
		}
	}
	for i, city := range expectedCities {
		if l := results.Items[i]; l != city {
			t.Errorf("肯定包含的用户 # %d: %s, 但计算得到的是: %s", i, city, l)
		}
	}

}
