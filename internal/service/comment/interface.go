package comment

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.Comment, error)
	Update(ctx context.Context, data Update) (entity.Comment, error)
	GetAll(ctx context.Context, filter Filter) ([]entity.Comment,int, error)
	GetTaskId(ctx context.Context, id int) ([]entity.Comment,int, error)
	Delete(ctx context.Context, id int, userId int) error
}
