package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {

	var op = zap.Fields(zap.String("LevelKey", "xxx"))
	// 可以自定义个别参数
	logger := zap.NewExample(op)

	// 使用sugar包装 性能不如 zap
	sugar := logger.Sugar()
	url := "www.xxx.com"
	sugar.Infow("链接失败", "url", url, "attempt", 3, "backoff", time.Second)
	sugar.Infof("链接失败 URL %s", url)
	// result
	//{"level":"info","msg":"链接失败","url":"www.xxx.com","attempt":3,"backoff":"1s"}
	//{"level":"info","msg":"链接失败 URL www.xxx.com"}
}
