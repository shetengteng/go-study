package main

import (
	"learngo/crawler_distributed/persist"
	"learngo/crawler_distributed/rpcsupport"

	"log"

	"fmt"
	"learngo/crawler_distributed/config"

	"flag"

	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for me to listen on")

// 开启一个rpc服务，该服务用于实现itemsaver功能
func main() {

	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		*port = config.ItemSaverPort
		//return
	}

	// 出错直接退出
	log.Fatal(StartServerRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex,
	))
}

// 开启服务rpc
// 第一参数是host，第二个参数是es的index
func StartServerRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}

	// 注册了ItemSaverService.Save方法，并开启服务
	err = rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
	return err
}
