package graphql

// ErrorCode are codes for representing errors in the service
type ErrorCode string

// ...
const (
	InputValidationError   ErrorCode = "InputValidationError"
	InternalSystemError ErrorCode = "InternalSystemError"
)

func (code *ErrorCode) String() string {
	return string(*code)
}
