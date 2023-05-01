package attendance

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/attendance"
)

type Attendance interface {
	Create(ctx context.Context, data attendance.Create) (entity.Attendance, error)
	GetById(ctx context.Context, id int) (entity.Attendance, error)
	GetAll(ctx context.Context, filter attendance.Filter) ([]entity.Attendance, error)
	Update(ctx context.Context, data attendance.Update) (entity.Attendance, error)
	Delete(ctx context.Context, id int, userId int) error
}
