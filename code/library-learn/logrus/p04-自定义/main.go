package main

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	// 转换为json格式
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Info("info msg")
}
