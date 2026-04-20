package logger

import (
	"context"
	"fmt"
	"runtime"
)

func GetRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(RequestIDKey).(string); ok {
		return v
	}
	return ""
}

// func GetCaller(ctx context.Context) string {
// 	if v, ok := ctx.Value(CallerKey).(string); ok {
// 		return v
// 	}
// 	return ""
// }

func getCaller() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	name := "unknown"
	if fn != nil {
		name = fn.Name()
	}

	return fmt.Sprintf("%s:%d %s", file, line, name)
}
