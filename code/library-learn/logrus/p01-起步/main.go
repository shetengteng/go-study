package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(false) // 设置输出文件名和方法信息

	// Trace最大 Panic最小
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	// 没有输出panic 由于在Fatal中进行了退出操作
	// panic 记录日志，然后panic
	logrus.Panic("panic msg")

	// result
	//time="2020-10-21T16:26:18+08:00" level=trace msg="trace msg"
	//time="2020-10-21T16:26:18+08:00" level=debug msg="debug msg"
	//time="2020-10-21T16:26:18+08:00" level=info msg="info msg"
	//time="2020-10-21T16:26:18+08:00" level=warning msg="warn msg"
	//time="2020-10-21T16:26:18+08:00" level=error msg="error msg"
	//time="2020-10-21T16:26:18+08:00" level=fatal msg="fatal msg"
}
