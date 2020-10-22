package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func main() {

	w1 := &bytes.Buffer{}
	w2 := os.Stdout
	w3, err := os.OpenFile("xx.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(w1, w2, w3))
	logrus.Info("info msg")
}
