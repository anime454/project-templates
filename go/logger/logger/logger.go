package logger

type LoggerPort interface {
	Debug(args ...any)
	// Info(args ...any)
	// Warn(args ...any)
	// Error(args ...any)

	// Debugf(format string, args ...any)
	// Infof(format string, args ...any)
	// Warnf(format string, args ...any)
	// Errorf(format string, args ...any)

	// With(fields any) Logger

	// Request(method, path string, fields any)
	// Response(status int, duration time.Duration, fields any)
	// Query(query string, duration time.Duration, fields any)
}
