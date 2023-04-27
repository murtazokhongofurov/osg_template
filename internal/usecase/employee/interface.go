package usecase

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/employee"
)

type Employee interface {
	Create(ctx context.Context, data employee.Create) (entity.Employee, error)
	GetById(ctx context.Context, id int) (entity.Employee, error)
	GetAll(ctx context.Context, filter employee.Filter) ([]entity.Employee, int, error)
	Update(ctx context.Context, data employee.Update) (entity.Employee, error)
	Delete(ctx context.Context, id int, userId int) error
}
