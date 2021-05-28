package errs

// PermissionError is used to represent an error showing that a session expired.
type PermissionError struct {
	Message string
}

// NewPermissionError returns a new NewPermissionError.
func NewPermissionError() PermissionError {
	return PermissionError{
		Message: "agent does not have the required permissions",
	}
}

func (err PermissionError) Error() string {
	return err.Message
}
