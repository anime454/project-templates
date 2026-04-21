package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	zl             zerolog.Logger
	maskingEnabled bool
	maskFields     map[string]any
}

func NewLogger(config LoggerConfig) LoggerPort {

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

func (l *Logger) ParseLogLevel(level string) LogLevel {
	return l.ParseLogLevel(level)
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

func (l *Logger) Request(requestLog HTTPRequestLog) {
	l.zl.Info().
		Str(FieldType, string(LogTypeRequest)).
		Str(FieldCaller, getCaller()).
		Str("request_timestamp", requestLog.Info.Timestamp.Format(time.RFC3339Nano)).
		Str("request_method", requestLog.Info.Method).
		Str("request_path", requestLog.Info.Path).
		Str("request_ip", requestLog.Info.IP).
		Str("request_protocol", requestLog.Info.Protocol).
		Str("http_request_id", requestLog.Meta.RequestID).
		Str("user_id", requestLog.Meta.UserID).
		Str("user_agent", requestLog.Meta.UserAgent).
		Interface("request_headers", l.maskValue(requestLog.Meta.Headers)).
		Interface("request_body", l.maskValue(requestLog.Body)).
		Send()
}

func (l *Logger) Response(responseLog HTTPResponseLog) {
	l.zl.Info().
		Str(FieldType, string(LogTypeResponse)).
		Str(FieldCaller, getCaller()).
		Str("response_timestamp", responseLog.Info.Timestamp.Format(time.RFC3339Nano)).
		Int("response_status", responseLog.Info.Status).
		Int64("response_size", responseLog.Info.Size).
		Str("response_protocol", responseLog.Info.Protocol).
		Str("http_request_id", responseLog.Meta.RequestID).
		Str("user_id", responseLog.Meta.UserID).
		Interface("response_headers", l.maskValue(responseLog.Meta.Headers)).
		Interface("response_body", l.maskValue(responseLog.Body)).
		Send()
}
