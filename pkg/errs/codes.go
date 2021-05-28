package errs

// MainErrorCode are codes for representing main errors in the services.
type MainErrorCode string

// ...
const (
	InputValidationError      MainErrorCode = "InputValidationError"
	PreSessionValidationError MainErrorCode = "PreSessionValidationError"
	SessionValidationError    MainErrorCode = "SessionValidationError"
	InvalidBasicAuth          MainErrorCode = "InvalidBasicAuth"
	InternalSystemError       MainErrorCode = "InternalSystemError"
)
