package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	// 注册接口方法
	rpc.Register(service)

	listener, error := net.Listen("tcp", host)

	if error != nil {
		return error
	}

	fmt.Println("ServeRpc ", host, " start")

	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Printf("accept error: %v", error)
			continue
		}
		// 处理连接的数据
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client := jsonrpc.NewClient(conn)
	return client, nil
}
