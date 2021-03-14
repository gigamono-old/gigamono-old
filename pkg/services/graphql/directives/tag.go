package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// Tag allows us to add custom struct tags to models.
func Tag(ctx context.Context, obj interface{}, next graphql.Resolver, rules *string) (interface{}, error) {
	return next(ctx)
}
