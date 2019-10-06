//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	"context"

	"github.com/rafaelrubbioli/espenses/generated"
	"github.com/rafaelrubbioli/espenses/models"
)

type Resolver struct{}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetExpenses(ctx context.Context, input models.GetExpensesInput) (
	*models.Expenses, error) {
	return nil, nil
}
