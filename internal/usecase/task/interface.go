package task

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/task"
)

type Task interface {
	Create(ctx context.Context, data task.Create) (entity.Task, error)
	Update(ctx context.Context, data task.Update) (entity.Task, error)
	GetById(ctx context.Context, id int) (entity.Task, error)
	GetAll(ctx context.Context, filter task.Filter) ([]entity.Task, int, error)
	Delete(ctx context.Context, id int, userId int) error
}
