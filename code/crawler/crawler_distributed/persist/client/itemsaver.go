package client

import (
	"learngo/crawler/engine"
	"log"

	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
)

// engine使用的saver客户端，调用远程save方法
func ItemSaver(host string) (chan engine.Item, error) {

	client, error := rpcsupport.NewClient(host)
	if error != nil {
		return nil, error
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("#%d got item %v ", itemCount, item)
			itemCount++

			// Call RPC to save
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Print("item saver : error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
