package errs

// ClientErrors are user-facing errors and they are mostly validation related.
// GQLGen already has a nice way of adding client errors.
// This is only relevant to REST endpoints.
type ClientErrors struct {
	ClientErrors []ClientError `json:"errors"`
}

// ClientError represents an error to be sent to the client.
type ClientError struct {
	Path            []string        `json:"path"`
	Message         string          `json:"message"`
	Code            MainErrorCode   `json:"code"`
	Type            ClientErrorType `json:"type,omitempty"`
}

// ClientErrorType is type of client error (body, params, query, etc.)
type ClientErrorType string

// ...
const (
	Query               ClientErrorType = "Query"
	Body                ClientErrorType = "Body"
	URLEncodedForm      ClientErrorType = "URLEncodedForm"
	Header              ClientErrorType = "Header"
	Cookie              ClientErrorType = "Cookie"
	ClientErrorTypeNone ClientErrorType = ""
)

func (code *ClientErrorType) String() string {
	return string(*code)
}

// AddError adds new client error to the bunch.
func (errs *ClientErrors) AddError(clientError ClientError) {
	errs.ClientErrors = append(errs.ClientErrors, clientError)
}

func (err ClientError) Error() string  {
	return err.Message
}
