package service

import (
	"context"

	"github.com/rafaelrubbioli/espenses/pkg/entity"
	"github.com/rafaelrubbioli/espenses/pkg/service/internal/repository"
)

type User interface{
	Create(ctx context.Context, name string) (int64, error)
	Get(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
}

type user struct {
	userRepository repository.User
}

func (u user) Create(ctx context.Context, name string) (int64, error) {
	return u.userRepository.Create(ctx, name)
}

func (u user) Get(ctx context.Context, id int64) (*entity.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u user) Update(ctx context.Context, user *entity.User) error {
	return u.userRepository.Update(ctx, user)
}

func (u user) Delete(ctx context.Context, id int64) error {
	return u.userRepository.Delete(ctx, id)
}

func NewUserService(userRepository repository.User) User {
	return user{
		userRepository: userRepository,
	}
}
