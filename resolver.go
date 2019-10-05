//go:generate go run github.com/99designs/gqlgen

package expenses

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetExpenses(ctx context.Context, input GetExpensesInput) (*Expenses, error) {
	panic("not implemented")
}
