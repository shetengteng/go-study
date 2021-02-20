package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"sword-man/src/core/constant"
	"time"
)

func init() {
	//registerInit(logConfig)
}

var (
	logConfig = new(LogConfig)
	loggerOut = new(lumberjack.Logger)
)

type LogConfig struct {
	Log struct {
		Level        string `validate:"oneof=trace debug info warn error fatal panic"`
		Formatter    string `validate:"oneof=json text nested"`
		Filename     string
		MaxSize      int // MB
		MaxBackups   int
		MaxAge       int  // 最多保留多少day，没有则一直保留
		Compress     bool // 是否压缩
		LocalTime    bool // 是否使用时间戳命名 backup 日志, 默认使用 UTC 格式
		EnableColors bool
	}
}

func (c *LogConfig) setupOrder() int {
	return 0
}

func (c *LogConfig) prepareSetup(v *viper.Viper) {
	viper.SetDefault("log.Level", "info")
	viper.SetDefault("log.Formatter", "nested")
	viper.SetDefault("log.Filename", constant.DEFAULT_LOG_PATH)
	viper.SetDefault("log.MaxSize", 500)
	viper.SetDefault("log.MaxBackups", -1)
	viper.SetDefault("log.MaxAge", -1)
	viper.SetDefault("log.Compress", false)
	viper.SetDefault("log.LocalTime", false)
	viper.SetDefault("log.EnableColors", false)
}

func (c *LogConfig) setupHandle() error {
	errorHandle(validator.New().Struct(c))

	level, _ := logrus.ParseLevel(c.Log.Level)
	logrus.SetLevel(level)

	switch c.Log.Formatter {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "nested":
		logrus.SetFormatter(&nested.Formatter{
			TimestampFormat: time.RFC3339,
			HideKeys:        false,
			NoColors:        !c.Log.EnableColors,
		})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat:  time.RFC3339,
			DisableTimestamp: false,
			DisableColors:    !c.Log.EnableColors,
			FullTimestamp:    true,
		})
	}

	loggerOut.Filename = c.Log.Filename
	loggerOut.MaxSize = c.Log.MaxSize
	loggerOut.Compress = c.Log.Compress
	loggerOut.LocalTime = c.Log.LocalTime

	if c.Log.MaxBackups != -1 {
		loggerOut.MaxBackups = c.Log.MaxBackups
	}
	if c.Log.MaxAge != -1 {
		loggerOut.MaxAge = c.Log.MaxAge
	}
	logrus.SetOutput(io.MultiWriter(loggerOut, os.Stdout))
	return nil
}

func FetchInitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
		HideKeys:        false,
		NoColors:        true,
	})
	loggerOut.Filename = constant.DEFAULT_LOG_PATH
	logger.SetOutput(io.MultiWriter(loggerOut, os.Stdout))
	return logger
}
