package project

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/project"
)

type Prject interface {
	Create(ctx context.Context, data project.Create) (entity.Project, error)
	GetById(ctx context.Context, id int) (entity.Project, error)
	GetAll(ctx context.Context, filter project.Filter) ([]entity.Project, int, error)
	Update(ctx context.Context, data project.Update) (entity.Project, error)
	Delete(ctx context.Context, id int, userId int) error
}
