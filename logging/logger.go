package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initializes the logger
func Init(lvl zerolog.Level) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).Level(lvl)
	log.Debug().Msg("Logger initialized")
}

func InitFromEnv() {

	var zlevel LogLevel

	level := os.Getenv("LOG_LEVEL")

	err := zlevel.Decode(level)

	if err == nil && len(level) > 0 {

		Init(zlevel.AsZeroLogLevel())

	} else {

		Init(zerolog.InfoLevel)
	}
}
