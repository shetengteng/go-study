package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	// step:
	// start itemSaverServer
	go StartServerRpc(":1234", "test1")

	// 先与es建立连接成功后在进行save操作
	time.Sleep(time.Second * 2)

	// start itemSaverClient
	client, err := rpcsupport.NewClient(":1234")
	if err != nil {
		panic(err)
	}

	// call save
	item := engine.Item{
		Url:  "http://www.baidu.com",
		Type: "zhenai",
		Id:   "100082",
		Payload: model.Profile{
			Age:        11,
			Height:     111,
			Weight:     22,
			Income:     "333-33332",
			Gender:     "女",
			Name:       "安静的",
			Xinzuo:     "射手座",
			Occupation: "人事/行政",
			Marriage:   "未婚",
			House:      "已购房",
			Hokou:      "山东",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s ;err %s", result, err)
	}

}
