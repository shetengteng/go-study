package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {

	logger := zap.NewExample() // 适合测试的时候使用
	//logger = zap.NewDevelopment() // dev
	//logger = zap.NewProduction() // prod
	defer logger.Sync() // 底层api有缓存，使用logger.Sync是同步到文件中

	// fmt.Printf 使用大量interface{}反射 而 zap 使用确定类型包装，如zap.String zap.Int等

	// 记录一条日志
	logger.Debug("连接失败", zap.String("url", "www.xxx.com"), zap.Int("attempt", 11), zap.Duration("backoff", time.Second))
	// result 默认是json格式
	// {"level":"info","msg":"连接失败","url":"www.xxx.com","attempt":11,"backoff":"1s"}
	logger.Info("连接失败", zap.String("url", "www.xxx.com"), zap.Int("attempt", 11), zap.Duration("backoff", time.Second))
	logger.Warn("连接失败", zap.String("url", "www.xxx.com"), zap.Int("attempt", 11), zap.Duration("backoff", time.Second))
	logger.Error("连接失败", zap.String("url", "www.xxx.com"), zap.Int("attempt", 11), zap.Duration("backoff", time.Second))

}
