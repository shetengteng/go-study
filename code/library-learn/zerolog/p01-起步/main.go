package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Print("hello world")
	// result
	// {"level":"debug","time":"2020-10-21T14:01:52+08:00","message":"hello world"}
	log.Debug().Str("extend", "xx").Msg("异常出现")
	// {"level":"debug","extend":"xx","time":"2020-10-21T14:03:47+08:00","message":"异常出现"}
	log.Debug().Str("name", "ss").Send()
	// {"level":"debug","name":"ss","time":"2020-10-21T14:04:16+08:00"}
	// 调用 Msg 或者 Send 之后，日志被打印

	// 嵌套 {"level":"info","dict":{"bar":"baz","n":1},"time":"2020-10-21T14:07:03+08:00","message":"hello world"}
	log.Info().Dict("dict", zerolog.Dict().
		Str("bar", "baz").
		Int("n", 1)).Msg("hello world")
}
