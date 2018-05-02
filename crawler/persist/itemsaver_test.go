package persist

import (
	"testing"
	"dali.cc/ccmouse/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"dali.cc/ccmouse/crawler/engine"
)

func TestSave(t *testing.T) {
	url := "http://album.zhenai.com/u/1077868794"
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
	// 可以用 docker go client去启一个独立的服务进行测试吗？ 不是在真实的服务中
	client, err := elastic.NewClient(elastic.SetSniff(false))
	const el_index = "profiles_test"
	err = Save(client,el_index, right)
	// todo 这里有一个依赖，第三方的服务。

	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index(el_index).
		Type(right.Type).Id(right.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual = engine.Item{}
	err = json.Unmarshal(*resp.Source, &actual)
	actual.Payload = model.Map2Profile(actual.Payload)

	if err != nil {
		panic(err)
	}
	if actual != right {
		t.Errorf("存的数据和取的数据不一致， \n %v \n %v", actual, right)
	}
}
