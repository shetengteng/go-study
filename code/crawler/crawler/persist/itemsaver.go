package persist

import (
	"log"

	"context"

	"learngo/crawler/engine"

	"errors"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(
		// 无法访问es内网，使用false，默认9200访问
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("#%d got item %v ", itemCount, item)
			itemCount++
			_, err := Save(client, item, index)
			if err != nil {
				log.Print("item saver : error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// 保存成功返回id
func Save(client *elastic.Client, item engine.Item, index string) (id string, err error) {

	if item.Type == "" {
		return "", errors.New("must supply Type")
	}

	indexSerivce := client.Index()
	if item.Id != "" {
		indexSerivce.Id(item.Id)
	}

	// 第一是Index操作，创建索引
	resp, err := indexSerivce.
		Index(index).
		Type(item.Type).
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	// +v可以打印详细字段名称
	//fmt.Printf("%+v", resp)
	return resp.Id, nil
}
