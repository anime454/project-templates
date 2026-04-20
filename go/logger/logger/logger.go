package logger

import (
	"context"

	"github.com/anime454/project-templates/go/logger/model"
)

type LoggerPort interface {
	Debug(arg any)
	Debugf(format string, args ...any)

	Info(arg any)
	Infof(format string, args ...any)

	Warn(msg string)
	Warnf(format string, args ...any)

	Error(err error)
	Errorf(format string, args ...any)

	With() *Logger
	WithContext(ctx context.Context) *Logger

	Request(httpRequest model.HTTPRequestLog)
	Response(httpResponse model.HTTPResponseLog)
	// Query(query string, duration time.Duration, fields any)
}
