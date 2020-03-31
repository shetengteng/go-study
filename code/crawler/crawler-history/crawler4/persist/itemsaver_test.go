package persist

import (
	"context"
	"testing"

	"encoding/json"

	"learngo/crawler-history/crawler4history/crawler4/engine"
	"learngo/crawler-history/crawler4history/crawler4/model"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
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

	// 获取es中的该数据，判断是一致
	// TODO try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	// &{Index:dating_profile Type:zhenai Id:AXElFUfWDzgmNPwEevfu Version:1 Result:created Shards:0xc042144720
	// SeqNo:0 PrimaryTerm:0 Status:0 ForcedRefresh:false Created:true}
	_, err = Save(client, expected, "dating_profile")
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	//t.Logf("%+v", resp)
	// source中存储了对象的json
	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v expected %v", actual, expected)
	}
}
