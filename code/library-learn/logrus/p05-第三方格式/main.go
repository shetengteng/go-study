package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        false, // 是否隐藏key，针对自定义的field
		TimestampFormat: time.RFC3339,
		FieldsOrder:     []string{"age", "name"}, // 排序 针对自定义的field
	})

	log := logrus.WithFields(logrus.Fields{
		"name": "ss",
		"age":  11,
	})

	// 2020-10-21T16:58:46+08:00 [INFO] [age:11] [name:ss] info msg
	log.Info("info msg")
	logrus.WithFields(logrus.Fields{"x":"y"}).Info("ss")
}
