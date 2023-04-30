package attendance

import (
	"context"
	"time"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/attendance"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewAttendanceRepo(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data attendance.Create) (entity.Attendance, error) {
	var detail entity.Attendance
	now := time.Now()
	detail.EmployeeId = data.EmployeeId
	detail.Type = data.Type
	detail.CreatedAt = now
	_, err := r.DB.NewInsert().Model(&detail).Exec(ctx)
	return detail, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.Attendance, error) {
	var detail entity.Attendance
	err := r.DB.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)
	return detail, err
}

func (r Repository) Update(ctx context.Context, data attendance.Update) (entity.Attendance, error) {
	var detail entity.Attendance
	detail.Id = data.Id
	if data.Type != nil {
		detail.Type = data.Type
	}
	_, err := r.DB.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)
	return detail, err
}

func (r Repository) GetAll(ctx context.Context, filter attendance.Filter) ([]entity.Attendance, int, error) {
	var list []entity.Attendance
	q := r.DB.NewSelect().Model(&list)
	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}
	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}
	count, err := q.ScanAndCount(ctx)
	return list, count, err
}

func (r Repository) Delete(ctx context.Context, id int, userId int) error {
	_, err := r.DB.NewUpdate().Table("attendance").
		Set("deleted_by = ?", userId).
		Set("deleted_at = ?", time.Now()).Where("id = ?", id).Exec(ctx)
	return err
}
