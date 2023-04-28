package developer

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Service struct {
	repo Repository
}

func NewDeveloperService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, data Create) (entity.Developer, error) {
	return s.repo.Create(ctx, data)
}

func (s Service) GetById(ctx context.Context, id int) (entity.Developer, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) Update(ctx context.Context, data Update) (entity.Developer, error) {
	return s.repo.Update(ctx, data)
}

func (s Service) GetAll(ctx context.Context, filter Filter) ([]entity.Developer, int,error) {
	return s.repo.GetAll(ctx, filter)
}

func (s Service) Delete(ctx context.Context, id int, userId int) error {
	return s.repo.Delete(ctx, id, userId)
}
