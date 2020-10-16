//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	"github.com/rafaelrubbioli/espenses/pkg/graphql/internal"
)

type Resolver struct{}

func (r *Resolver) Query() internal.QueryResolver {
	return &query{r}
}
