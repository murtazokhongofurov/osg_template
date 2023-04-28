package comment

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Service struct {
	repo Repository
}

func NewCommentService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, data Create) (entity.Comment, error) {
	return s.repo.Create(ctx, data)
}

func (s Service) GetAll(ctx context.Context, filter Filter) ([]entity.Comment, int, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s Service) GetTaskId(ctx context.Context, id int) ([]entity.Comment, int, error) {
	return s.repo.GetTaskId(ctx, id)
}

func (s Service) Update(ctx context.Context, data Update) (entity.Comment, error) {
	return s.repo.Update(ctx, data)
}

func (s Service) Delete(ctx context.Context, id int, userId int) error {
	return s.repo.Delete(ctx, id, userId)
}
