package resolver

import (
	"context"

	"github.com/rafaelrubbioli/espenses/pkg/graphql/internal/models"
)

type mutation struct{ *resolver }

func (m mutation) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	id, err := m.services.UserService.Create(ctx, input.Name, input.Password)
	if err != nil {
		return nil, err
	}

	user, err := m.services.UserService.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.NewUser(*user), nil
}
