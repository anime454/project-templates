package db

import "fmt"

type DBError struct {
	Code    string
	Message string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewDBError(code, message string) *DBError {
	return &DBError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrRecordNotFound *DBError = NewDBError("DB_RECORD_NOT_FOUND", "record not found")
)
