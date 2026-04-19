package logger

import (
	"context"
	"os"
	"time"

	"github.com/anime454/project-templates/go/logger/model"
	"github.com/rs/zerolog"
)

const defaultMaskValue = "******"

type Logger struct {
	zl             zerolog.Logger
	maskingEnabled bool
	maskFields     map[string]any
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

	return &Logger{
		zl:             zl,
		maskingEnabled: config.Masking.Enabled,
		maskFields:     normalizeMaskFields(config.Masking.FieldMap),
	}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	reqID := GetRequestID(ctx)
	caller := getCallerOfWithContext()

	newLogger := l.zl.With().
		Str(string(RequestIDKey), reqID).
		Str(string(CallerKey), caller).
		Logger()

	return &Logger{zl: newLogger, maskingEnabled: l.maskingEnabled, maskFields: l.maskFields}
}

func (l *Logger) Debug(arg any) {
	l.zl.Debug().
		Interface(MessageField, l.maskValue(arg)).
		Str(FieldType, string(LogTypeDebug)).
		Send()
}

func (l *Logger) Info(arg any) {
	l.zl.Info().
		Interface(MessageField, l.maskValue(arg)).
		Str(FieldType, string(LogTypeInfo)).
		Send()
}

func (l *Logger) Debugf(format string, args ...any) {
	l.zl.Debug().Msgf(format, l.maskArgs(args)...)
}
