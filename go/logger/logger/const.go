package logger

const (
	MESSAGE_FIELD = "message"
	// REQUEST_ID_FIELD = "request_id"
	// DEBUG_LOG_FIELD  = "debug"
)

type contextKey string

const RequestIDKey contextKey = "request_id"
const CallerKey contextKey = "caller"
