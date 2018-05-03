package main

import (
	"testing"
	"dali.cc/ccmouse/crawler_distributed/rpcsupport"
	"dali.cc/ccmouse/crawler/engine"
	"dali.cc/ccmouse/crawler/model"
	"time"
)

// 启动服务
// 调用服务
// 查看结果
func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test_profile")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
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
	result := ""
	client.Call("ItemSaverService.Save", right, &result)
	if err != nil || result != "ok" {
		t.Errorf("result is %s, err is %s", result, err)
	}
	// todo 从 elastic 验证测试数据
}
