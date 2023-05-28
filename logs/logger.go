package logs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func customLogger() zerolog.Logger {
	logFile, err := os.OpenFile(
		"./logs/log-file.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)

	if err != nil {
		log.Panic().Msgf("Could not open file due to: %s", err.Error())
	}

	logger := zerolog.New(logFile).With().Timestamp().Caller().Logger()
	return logger
}

var Logger = customLogger()
