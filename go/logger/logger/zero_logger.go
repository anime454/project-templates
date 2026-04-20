package logger

import (
	"context"
	"os"
	"time"

	"github.com/anime454/project-templates/go/logger/model"
	"github.com/rs/zerolog"
)

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
	zerolog.ErrorFieldName = FieldMessage

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

func (l *Logger) With() *Logger {
	return &Logger{zl: l.zl.With().Logger(), maskingEnabled: l.maskingEnabled, maskFields: l.maskFields}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	reqID := GetRequestID(ctx)

	newLogger := l.zl.With().
		Str(string(RequestIDKey), reqID).
		Logger()

	return &Logger{zl: newLogger, maskingEnabled: l.maskingEnabled, maskFields: l.maskFields}
}

func (l *Logger) Debug(arg any) {
	l.zl.Debug().
		Interface(FieldMessage, l.maskValue(arg)).
		Str(FieldType, string(LogTypeDebug)).
		Str(FieldCaller, getCaller()).
		Send()
}

func (l *Logger) Debugf(format string, args ...any) {
	l.zl.Debug().
		Str(FieldType, string(LogTypeDebug)).
		Str(FieldCaller, getCaller()).
		Msgf(format, l.maskArgs(args)...)
}

func (l *Logger) Info(arg any) {
	l.zl.Info().
		Interface(FieldMessage, l.maskValue(arg)).
		Str(FieldType, string(LogTypeInfo)).
		Str(FieldCaller, getCaller()).
		Send()
}

func (l *Logger) Infof(format string, args ...any) {
	l.zl.Info().
		Str(FieldType, string(LogTypeInfo)).
		Str(FieldCaller, getCaller()).
		Msgf(format, l.maskArgs(args)...)
}

func (l *Logger) Warn(msg string) {
	l.zl.Warn().
		Str(FieldMessage, msg).
		Str(FieldType, string(LogTypeWarn)).
		Str(FieldCaller, getCaller()).
		Send()
}

func (l *Logger) Warnf(format string, args ...any) {
	l.zl.Warn().
		Str(FieldType, string(LogTypeWarn)).
		Str(FieldCaller, getCaller()).
		Msgf(format, l.maskArgs(args)...)
}

func (l *Logger) Error(err error) {
	l.zl.Error().
		Err(err).
		Str(FieldType, string(LogTypeError)).
		Str(FieldCaller, getCaller()).
		Send()
}

func (l *Logger) Errorf(format string, args ...any) {
	l.zl.Error().
		Str(FieldType, string(LogTypeError)).
		Str(FieldCaller, getCaller()).
		Msgf(format, l.maskArgs(args)...)
}
