package main

import "github.com/sirupsen/logrus"

func main() {

	// 添加公共字段，用于识别人员或机器信息等
	log := logrus.WithFields(logrus.Fields{
		"userId": 111,
		"IP":     "111.111.111.111",
	})
	log.Info("info msg")
}
