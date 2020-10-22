package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

func main() {
	// 提供了一个ConsoleWriter可输出便于我们阅读的，带颜色的日志。调用zerolog.Output()来启用ConsoleWriter：
	//logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	//logger.Info().Str("foo", "bar").Msg("hello world")

	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	// 2020-10-21T14:23:24+08:00 | INFO  | ***hello world**** foo:BAR
	logger := log.Output(output).With().Timestamp().Logger()
	logger.Info().Str("foo", "bar").Msg("hello world")
	// 由于每次格式需要钩子函数做转换，因此性能有影响，不建议在生产中使用

}
