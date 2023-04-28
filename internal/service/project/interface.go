package project

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.Project, error)
	GetById(ctx context.Context, id int) (entity.Project, error)
	GetAll(ctx context.Context, filter Filter) ([]entity.Project, int, error)
	Update(ctx context.Context, data Update) (entity.Project, error)
	Delete(ctx context.Context, id int, userId int) error
}
