package services

type BusinessError struct {
	Code    string
	Message string
}

func (e *BusinessError) Error() string {
	return e.Code + ": " + e.Message
}

func NewBusinessError(code, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrVehicleAlreadyCheckedIn *BusinessError = NewBusinessError("VEHICLE_001", "vehicle is already checked in")
)
