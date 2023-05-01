package task

import (
	"context"
	"errors"

	"github.com/osg_template/internal/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, data Create) (entity.Task, error) {
	if data.Title == nil {
		return entity.Task{}, errors.New("title is not required")
	}
	if data.Description == nil {
		return entity.Task{}, errors.New("description is not required")
	}

	return s.repo.Create(ctx, data)
}

func (s Service) Update(ctx context.Context, data Update) (entity.Task, error) {
	return s.repo.Update(ctx, data)
}

func (s Service) GetAll(ctx context.Context, filter Filter) ([]entity.Task, int, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s Service) GetById(ctx context.Context, id int) (entity.Task, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) Delete(ctx context.Context, id int, userId int) error {
	return s.repo.Delete(ctx, id, userId)
}
