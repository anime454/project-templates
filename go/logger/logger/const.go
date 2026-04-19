package logger

const (
	MessageField = "message"
	FieldType    = "type"
)

type LogType string

const (
	LogTypeDebug    LogType = "debug"
	LogTypeInfo     LogType = "info"
	LogTypeRequest  LogType = "request"
	LogTypeResponse LogType = "response"
	LogTypeQuery    LogType = "query"
)

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	CallerKey    contextKey = "caller"
)
