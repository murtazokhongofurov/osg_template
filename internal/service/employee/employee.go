package employee

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

func (s Service) Create(ctx context.Context, data Create) (entity.Employee, error) {
	if data.FullName == nil {
		return entity.Employee{}, errors.New("fullname is not required")
	}
	if data.Role == nil {
		return entity.Employee{}, errors.New("role is not required")
	}
	return s.repo.Create(ctx, data)
}

func (s Service) GetById(ctx context.Context, id int) (entity.Employee, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) GetAll(ctx context.Context, filter Filter) ([]entity.Employee, int, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s Service) Update(ctx context.Context, data Update) (entity.Employee, error) {
	return s.repo.Update(ctx, data)
}

func (s Service) Delete(ctx context.Context, id int, userId int) error {
	return s.repo.Delete(ctx, id, userId)
}
