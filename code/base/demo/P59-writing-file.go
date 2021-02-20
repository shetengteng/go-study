package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 直接创建并写入内容到文件
	// 注意：运行前先创建tmp文件夹
	//
	//os.MkdirAll(path.Dir(fn))
	//... = os.Create(fn)

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("d:/tmp/dat1", d1, 0644)
	check2(err)

	// 打开文件，然后开始写
	f, err := os.Create("d:/tmp/dat2")
	check2(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check2(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check2(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// flush 数据到本地存储中
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check2(err)
	fmt.Printf("wrote %d bytes \n", n4)

	w.Flush()
}
