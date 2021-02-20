package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 读取全部文件到内存中
	dat, err := ioutil.ReadFile("tmp/dat")
	check(err)
	fmt.Println(string(dat))

	// 如果是要读取文件的一部分，则先打开文件
	f, err := os.Open("tmp/dat")
	check(err)

	// 读取开始部分
	b1 := make([]byte, 5)
	// 返回的n1表示读取的字节数目
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 读取中间部分，先跳转到该位置
	o2, err := f.Seek(6, 0)
	check(err)

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	// 从o2的位置开始读取 n2个字节
	fmt.Printf("%d bytes @ %d", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// 如果读的数目没有满足最小的要求的数目的数据，则抛出异常ErrUnexpectedEOF
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 回退到开头
	_, err = f.Seek(0, 0)
	check(err)

	// 缓冲流
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
