package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// 使用filepath包的工具进行路径的生成，不同系统的路径格式不同
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p", p)

	// 不建议手动连接路径，使用filepath.join，会去除冗余格式，规范化路径
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// dir 和 base 方法会自动分离出路径和文件
	fmt.Println("Dir(p):",filepath.Dir(p))
	fmt.Println("Base(p):",filepath.Base(p))

	filename := "config.json"
	// 分离出扩展名，含有.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 可以分离出文件名，不带扩展名
	fmt.Println(strings.TrimSuffix(filename,ext))

	//
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
