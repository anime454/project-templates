package logger

import "context"

type LoggerPort interface {
	Debug(arg any)
	Info(arg any)
	// Warn(msg string)
	// Error(msg string)

	Debugf(format string, args ...any)
	// Infof(format string, args ...any)
	// Warnf(format string, args ...any)
	// Errorf(format string, args ...any)

	WithContext(ctx context.Context) *Logger

	// With(fields any) Logger

	// Request(method, path string, fields any)
	// Response(status int, duration time.Duration, fields any)
	// Query(query string, duration time.Duration, fields any)
}
