package comment

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/comment"
)

type Comment interface {
	Create(ctx context.Context, data comment.Create) (entity.Comment, error)
	GetTaskId(ctx context.Context, id int) ([]entity.Comment, int, error)
	GetAll(ctx context.Context, filter comment.Filter) ([]entity.Comment, int, error)
	Update(ctx context.Context, data comment.Update) (entity.Comment, error)
	Delete(ctx context.Context, id int, userId int) error
}
