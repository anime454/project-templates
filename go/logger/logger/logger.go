package logger

import (
	"context"
)

type LoggerPort interface {
	With() *Logger
	WithContext(ctx context.Context) *Logger

	Debug(arg any)
	Debugf(format string, args ...any)

	Info(arg any)
	Infof(format string, args ...any)

	Warn(msg string)
	Warnf(format string, args ...any)

	Error(err error)
	Errorf(format string, args ...any)

	Request(httpRequest HTTPRequestLog)
	Response(httpResponse HTTPResponseLog)
	// Query(query string, duration time.Duration, fields any)

}
