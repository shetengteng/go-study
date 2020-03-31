package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

// 转二进制
func convertToBin(n int) string {
	result := ""
	for; n> 0;n >>=1  {
		lsb:= n % 2
		result = strconv.Itoa(lsb)+result
	}
	return result;
}

func printFile(filename string){
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	fmt.Println(convertToBin(11))
	printFile("abc.txt")
}
