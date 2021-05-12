package response

import "github.com/gigamono/gigamono/pkg/errs"

// Response is a REST endpoint response.
type Response struct {
	Errors  []errs.ClientError `json:"errors,omitempty"`
	Message string             `json:"message,omitempty"`
	Data    interface{}        `json:"data,omitempty"`
}
