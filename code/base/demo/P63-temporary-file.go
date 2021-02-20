package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)


func check4(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// 如果第一个参数为空，则路径为os.TempDir()
	// 第二个参数是文件名称的前缀，每次创建的文件名不同，防止并发情况下冲突
	f, err := ioutil.TempFile("", "sample")
	check4(err)
	// 在linux中，临时文件一般在tmp文件夹下
	fmt.Println("temp file name:",f.Name())
	defer os.RemoveAll(f.Name())

	_, err = f.Write([]byte{1, 2, 3})
	check4(err)

	// 创建临时文件夹
	dname, err := ioutil.TempDir("", "sampledir")
	check4(err)
	fmt.Println("Temp dir name:",dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check4(err)
}
