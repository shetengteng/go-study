package main

import (
	"fmt"
	"log"
	"net/http"
	"stt"
)

// 模拟数据库
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	stt.NewGroup("scores", 2<<10, stt.GetterFunc(
		// 模拟数据库查询
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:9999"
	peers := stt.NewHTTPPool(addr)
	log.Println("stt is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
