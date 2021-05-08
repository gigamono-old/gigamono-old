package response

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Error returns a client error.
func Error(ctx context.Context, message string, otherPath ...ast.PathElement) error {
	rootPath := graphql.GetPath(ctx)
	path := append(rootPath, otherPath...)
	return &gqlerror.Error{
		Path:    path,
		Message: message,
		Extensions: map[string]interface{}{
			"code": errs.InputValidationError,
		},
	}
}
