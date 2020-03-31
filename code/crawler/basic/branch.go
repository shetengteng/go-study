package main

import (
	"io/ioutil"
	"fmt"
)

const filename = "abc.txt"



func read1(){
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
}

func read2(){
	if contents, err := ioutil.ReadFile(filename);
	err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
}

func grade(score int) string {
	g := ""
	switch  {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong score: %d",score))
	}
	return g
}

func main() {
	read1()
	read2()
	fmt.Println(
		grade(0),
		grade(10),
		grade(100),
		grade(100),
	)
}