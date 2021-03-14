package interceptors

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sageflow/sageflow/pkg/logs"
	"github.com/vektah/gqlparser/v2/gqlerror"
	gql "github.com/sageflow/sageflow/pkg/services/graphql"
)

type (
	// ErrorModifier intercepts error messages and modifies them accordingly.
	ErrorModifier struct{}
)

var _ interface {
	graphql.HandlerExtension
	graphql.ResponseInterceptor
} = ErrorModifier{}

// InterceptResponse modifies error messages with "internal system error".
func (interceptor ErrorModifier) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	response := next(ctx)

	// Check if error messages contain a "internal system error".
	errors := graphql.GetErrors(ctx)
	for _, err := range errors {
		if err.Message == "internal system error" {
			// Log response before changing.
			logs.FmtPrintf("intercepted server error:\n%v", response.Errors)

			// Change response to only show server error.
			response = &graphql.Response{
				Errors: gqlerror.List{{
					Message: "internal system error",
					Extensions: map[string]interface{}{
						"code": gql.InternalSystemError,
					},
				}},
			}

			break
		}
	}

	return response
}

// ExtensionName returns extension's name.
func (interceptor ErrorModifier) ExtensionName() string {
	return "ErrorResponseInterceptor"
}

// Validate an interface impl does nothing.
func (interceptor ErrorModifier) Validate(schema graphql.ExecutableSchema) error {
	return nil
}
