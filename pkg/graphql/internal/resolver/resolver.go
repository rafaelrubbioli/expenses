//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	"github.com/rafaelrubbioli/espenses/pkg/graphql/internal"
	"github.com/rafaelrubbioli/espenses/pkg/service"
)

type resolver struct{
	services service.All
}

func (r *resolver) Query() internal.QueryResolver {
	return &query{r}
}

func (r *resolver) Mutation() internal.MutationResolver {
	return &mutation{r}
}

func New(services service.All) internal.ResolverRoot {
	return &resolver{services: services}
}
