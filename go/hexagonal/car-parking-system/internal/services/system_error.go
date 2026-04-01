package services

type SystemError struct {
	Code    string
	Message string
}

func (e *SystemError) Error() string {
	return e.Code + ": " + e.Message
}

func NewSystemError(code, message string) *SystemError {
	return &SystemError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrInternalServerError *SystemError = NewSystemError("SYSTEM_9999", "internal server error")
	ErrDatabaseError       *SystemError = NewSystemError("SYSTEM_9998", "database error")
	ErrDatabaseQueryError  *SystemError = NewSystemError("SYSTEM_9997", "database query error")
)
