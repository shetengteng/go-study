package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// panic/fatal/error/warn/info/debug/trace
	debug := flag.Bool("debug", true, "sets log level to debug")
	flag.Parse()

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Debug().Msg("This message appears only when log level set to debug")
	log.Info().Msg("This message appears when log level set to debug or info")

	if e := log.Debug(); e.Enabled() {
		e.Str("foo", "bar").Msg("some debug message")
	}

}
