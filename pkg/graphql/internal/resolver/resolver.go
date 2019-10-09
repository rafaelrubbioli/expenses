//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	"context"

	"github.com/rafaelrubbioli/espenses/pkg/graphql/internal"
	"github.com/rafaelrubbioli/espenses/pkg/graphql/internal/models"
)

type Resolver struct{}

func (r *Resolver) Query() internal.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetExpenses(ctx context.Context, input models.GetExpensesInput) (
	*models.Expenses, error) {
	return &models.Expenses{}, nil
}
