package developer

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.Developer, error)
	GetById(ctx context.Context, id int) (entity.Developer, error)
	GetAll(ctx context.Context, filter Filter) ([]entity.Developer, int, error)
	Update(ctx context.Context, data Update) (entity.Developer, error)
	Delete(ctx context.Context, id int, userId int) error
}
