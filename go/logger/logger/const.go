package logger

type LogLevel int8

// Reference:
// http://github.com/rs/zerolog/blob/master/log.go#L127

const (
	// DebugLevel defines debug log level.
	DebugLevel LogLevel = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel LogLevel = -1
	// Values less than TraceLevel are handled as numbers.
)

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
