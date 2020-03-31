package main

import (
	"flag"
	"fmt"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
	"log"
)

// 命令行参数
// go run worker-starter.go --help 可以查看port的提示语句
// go run worker-starter.go --port=9000指定启动port
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {

	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{},
	))

}
