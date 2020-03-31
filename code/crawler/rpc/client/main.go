package main

import (
	"fmt"
	rpcdemo "learngo/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {

	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call(
		"DemoService.Div",
		rpcdemo.Args{10, 3},
		&result,
	)
	fmt.Println(result, err)

	var result2 float64
	err = client.Call(
		"DemoService.Div",
		rpcdemo.Args{10, 0},
		&result2,
	)
	fmt.Println(result2, err)
}
