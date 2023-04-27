package employee

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.Employee, error)
	GetById(ctx context.Context, id int) (entity.Employee, error)
	GetAll(ctx context.Context, filter Filter) ([]entity.Employee, int, error)
	Update(ctx context.Context, data Update) (entity.Employee, error)
	Delete(ctx context.Context, id int, userId int) error
}
