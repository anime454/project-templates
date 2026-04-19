package logger

import (
	"os"
	"time"

	"github.com/anime454/project-templates/go/logger/model"
	"github.com/rs/zerolog"
)

type Logger struct {
	zl *zerolog.Logger
}

func NewLogger(config model.LoggerConfig) LoggerPort {

	// set time format
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	// set logger level
	zerolog.SetGlobalLevel(zerolog.Level(config.Level))

	zl := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()

	return &Logger{zl: &zl}
}

func (l *Logger) Debug(args ...any) {
	l.zl.Info().Interface("args", args).Msg("")
}
