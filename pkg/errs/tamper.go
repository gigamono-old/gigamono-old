package errs

// TamperError is used to represent an error showing that a sensitive data has been tampered with or doesn't match what is expected.
type TamperError struct {
	Message string
}

// NewTamperError returns a new NewTamperError.
func NewTamperError() TamperError {
	return TamperError{
		Message: "values do not match or signature do not match",
	}
}

func (err TamperError) Error() string {
	return err.Message
}
