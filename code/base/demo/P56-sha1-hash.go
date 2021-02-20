package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))
	// sha1求和
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
