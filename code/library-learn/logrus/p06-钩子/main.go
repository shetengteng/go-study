package main

import "github.com/sirupsen/logrus"

// 钩子需要实现对应的接口 logrus.Hook
type MyHook struct {
	Name string
}

func (m *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels // 针对所有的日志都进行触发钩子操作
}

func (m *MyHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = m.Name // 在输出的数据中添加自定义字段
	return nil
}

func main() {
	h := &MyHook{Name: "stt"}
	logrus.AddHook(h)
	logrus.Info("info msg")
}
