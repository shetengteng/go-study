package main

import (
	"fmt"
	"log"
	"net/http"
	"stt"
)

func main() {

	router := stt.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
	})

	router.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = [%q]\n", k, v)
		}
	})

	log.Fatal(router.Run(":9999"))
}
