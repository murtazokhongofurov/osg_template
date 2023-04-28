package developer

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/developer"
)

type Developer interface {
	Create(ctx context.Context, data developer.Create) (entity.Developer, error)
	GetById(ctx context.Context, id int) (entity.Developer, error)
	GetAll(ctx context.Context, filter developer.Filter) ([]entity.Developer, int, error)
	Update(ctx context.Context, data developer.Update) (entity.Developer, error)
	Delete(ctx context.Context, id int, userId int) error
}
