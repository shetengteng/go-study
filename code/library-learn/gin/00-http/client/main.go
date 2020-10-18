package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:8090/go")

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)

	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		fmt.Println("read completed")
		res := string(buf[:n])
		fmt.Println(res)
		break
	}
	defer resp.Body.Close()
}
