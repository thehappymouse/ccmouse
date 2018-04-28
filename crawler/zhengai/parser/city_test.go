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
	results := ParseCity(body, "")

	var expectedUrls = []string{
		"http://album.zhenai.com/u/1861056016",
		"http://album.zhenai.com/u/106344107",
		"http://album.zhenai.com/u/1077868794",
	}

	for i, url := range expectedUrls {
		if l := results.Requests[i].Url; l != url {
			t.Errorf("肯定包含的url # %d: %s, 但计算得到的是: %s", i, url, l)
		}
	}
}
