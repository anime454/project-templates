package logger

const defaultMaskValue = "******"

const (
	FieldMessage = "message"
	FieldType    = "type"
	FieldCaller  = "caller"
)

type LogType string

const (
	LogTypeDebug    LogType = "debug"
	LogTypeInfo     LogType = "info"
	LogTypeRequest  LogType = "request"
	LogTypeResponse LogType = "response"
	LogTypeQuery    LogType = "query"
	LogTypeWarn     LogType = "warn"
	LogTypeError    LogType = "error"
)

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
)
