package resolver

import (
	"context"

	"github.com/rafaelrubbioli/espenses/pkg/graphql/internal/models"
)

type query struct{ *Resolver }

func (r *query) GetExpenses(ctx context.Context, input models.GetExpensesInput) (*models.Expenses, error) {
	return &models.Expenses{}, nil
}
