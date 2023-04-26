package task

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.Task, error)
	Update(ctx context.Context, data Update) (entity.Task, error)
	GetById(ctx context.Context, id int) (entity.Task, error)
	GetAll(ctx context.Context, filter Filter) ([]entity.Task, int, error)
	Delete(ctx context.Context, id int, userId int) error
}
