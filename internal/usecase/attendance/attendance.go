package attendance

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/attendance"
)

type UseCase struct {
	repo Attendance
}

func NewAttenUseCase(repo Attendance) *UseCase {
	return &UseCase{repo: repo}
}

func (au UseCase) CreateAttendance(ctx context.Context, data attendance.Create) (entity.Attendance, error) {
	return au.repo.Create(ctx, data)
}

func (au UseCase) DeleteAttendance(ctx context.Context, id int, userId int) error {
	return au.repo.Delete(ctx, id, userId)
}
