package graphql

import (
	"context"
	"errors"
	"strings"
	"unicode"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ValidateStructAndAppendErrors a struct and appends errors, if there is one, to response.
func ValidateStructAndAppendErrors(ctx context.Context, validate *validator.Validate, obj interface{}, argName string) error {
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	if validationErrors := err.(validator.ValidationErrors); validationErrors != nil {
		for _, err := range validationErrors {
			// Recreate the rule that failed.
			rule := err.ActualTag()
			param := err.Param()
			if param != "" {
				rule += "=" + param
			}

			// Recreate graphql path for input argument.
			path := append(graphql.GetPath(ctx), ast.PathName(argName))
			namespace := strings.Split(err.Namespace(), ".")[1:]
			for _, name := range namespace {
				path = append(path, ast.PathName(makeFirstCharLowercase(name)))
			}

			// Add error to response.
			graphql.AddError(ctx, &gqlerror.Error{
				Path:    path,
				Message: err.Error(),
				Extensions: map[string]interface{}{
					"code":  errs.InputValidationError,
					"value": err.Value(),
					"rule":  rule,
				},
			})
		}
	}

	return errors.New("")
}

func makeFirstCharLowercase(s string) string {
	for _, v := range s {
		return string(unicode.ToLower(v)) + s[1:]
	}

	return s
}
