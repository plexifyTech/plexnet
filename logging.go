package plexnet

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func InitLogging(config *AppConfig) {
	zerolog.DurationFieldInteger = true
	zerolog.CallerMarshalFunc = shortCallerMarshaller
	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Caller().
		Str("app", config.AppName).
		Logger()
}

func shortCallerMarshaller(pc uintptr, file string, line int) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}
