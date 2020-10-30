package service

import (
	"context"

	"github.com/rafaelrubbioli/espenses/pkg/entity"
	"github.com/rafaelrubbioli/espenses/pkg/service/internal/repository"
)

type Tag interface{
	Create(ctx context.Context, name string) (int64, error)
	Get(ctx context.Context, id int64) (*entity.Tag, error)
	Update(ctx context.Context, user *entity.Tag) error
	Delete(ctx context.Context, id int64) error
}

type tag struct {
	tagRepository repository.Tag
}

func (u tag) Create(ctx context.Context, name string) (int64, error) {
	return u.tagRepository.Create(ctx, name)
}

func (u tag) Get(ctx context.Context, id int64) (*entity.Tag, error) {
	return u.tagRepository.Get(ctx, id)
}

func (u tag) Update(ctx context.Context, tag *entity.Tag) error {
	return u.tagRepository.Update(ctx, tag)
}

func (u tag) Delete(ctx context.Context, id int64) error {
	return u.tagRepository.Delete(ctx, id)
}

func NewTagService(repo repository.Tag) Tag {
	return tag{
		tagRepository: repo,
	}
}
