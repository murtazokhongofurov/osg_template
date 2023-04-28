package attendance

import (
	"context"

	"github.com/osg_template/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.Attendance, error)
	GetById(ctx context.Context, id int) (entity.Attendance, error)
	GetAll(ctx context.Context, filter Filter) ([]entity.Attendance,int,  error)
	Update(ctx context.Context, data Update) (entity.Attendance, error)
	Delete(ctx context.Context, id int, userId int) error
}
