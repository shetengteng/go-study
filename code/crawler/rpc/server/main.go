package main

import (
	"learngo/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	rpc.Register(rpcdemo.DemoService{})

	listener, error := net.Listen("tcp", ":1234")

	if error != nil {
		panic(error)
	}

	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Printf("accept error: %v", error)
			continue
		}
		// 处理连接的数据
		go jsonrpc.ServeConn(conn)
	}
}
