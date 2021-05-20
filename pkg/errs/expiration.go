package errs

// ExpirationError is used to represent an error showing that a session expired.
type ExpirationError struct {
	Message string
}

// NewExpirationError returns a new NewExpirationError.
func NewExpirationError() ExpirationError {
	return ExpirationError{
		Message: "session expired",
	}
}

func (err ExpirationError) Error() string {
	return err.Message
}
