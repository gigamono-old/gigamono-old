package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type directiveFunc = func(context.Context, interface{}, graphql.Resolver, *string) (interface{}, error)
