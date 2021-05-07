package errs

// ErrorCode are codes for representing errors in the service
type ErrorCode string

// ...
const (
	InputValidationError ErrorCode = "InputValidationError"
	InternalSystemError  ErrorCode = "InternalSystemError"
	UnsupportedGrantType ErrorCode = "UnsupportedGrantType"
)

func (code *ErrorCode) String() string {
	return string(*code)
}
