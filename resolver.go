//go:generate gorunpkg github.com/99designs/gqlgen

package reqcheck

import (
	"context"

	"github.com/vektah/gqlgen/graphql"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetThing(ctx context.Context, id *string) (*Thing, error) {
	_ = graphql.CollectFieldsCtx(ctx, []string{"Thing"})

	return &Thing{}, nil
}
