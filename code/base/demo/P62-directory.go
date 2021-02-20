package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check3(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 创建相对文件夹，临时文件夹，在项目的根路径下
	err := os.Mkdir("subdir", 0755)
	check3(err)

	// 删除该文件夹下所有文件，类似rm -rf
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check3(ioutil.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")

	// 递归创建文件夹
	err = os.MkdirAll("subdir/parent/child", 0755)
	check3(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := ioutil.ReadDir("subdir/parent")
	check3(err)

	fmt.Println("listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 切换到指定文件夹下
	err = os.Chdir("subdir/parent/child")
	check3(err)

	// 读取当前文件夹下的所有文件与文件夹
	c, err = ioutil.ReadDir(".")
	check3(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 回退
	err = os.Chdir("../../..")
	check3(err)

	fmt.Println("Visiting subdir")
	// 传入一个回调方法，在遍历每个文件与文件夹的时候做出相应的动作
	err = filepath.Walk("subdir", visit)

	//fmt.Println(1/0)

}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
