package employee

import (
	"context"
	"time"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/employee"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data employee.Create) (entity.Employee, error) {
	var detail entity.Employee
	detail.FullName = data.FullName
	detail.ProfilePhoto = data.ProfilePhoto
	detail.Phone = data.Phone
	detail.BirthDate = data.BirthDate
	detail.Role = data.Role
	detail.Position = data.Position
	detail.CreatedBy = data.CreatedBy
	_, err := r.DB.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.Employee, error) {
	var detail entity.Employee
	err := r.DB.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)
	return detail, err
}

func (r Repository) GetAll(ctx context.Context, filter employee.Filter) ([]entity.Employee, int, error) {
	var list []entity.Employee
	q := r.DB.NewSelect().Model(&list)
	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}
	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}
	if filter.Search != nil {
		q.WhereGroup(" and ", func(sq *bun.SelectQuery) *bun.SelectQuery {
			sq.Where("full_name ilike ?", filter.Search)
			return sq
		})
	}
	count, err := q.ScanAndCount(ctx)
	return list, count, err
}

func (r Repository) Update(ctx context.Context, data employee.Update) (entity.Employee, error) {
	var detail entity.Employee
	detail.Id = data.Id

	if data.FullName != nil {
		detail.FullName = data.FullName
	}
	if data.ProfilePhoto != nil {
		detail.ProfilePhoto = data.ProfilePhoto
	}
	if data.BirthDate != nil {
		detail.BirthDate = data.BirthDate
	}
	if data.Phone != nil {
		detail.Phone = data.Phone
	}
	if data.Position != nil {
		detail.Position = data.Position
	}
	if data.Role != nil {
		detail.Role = data.Role
	}
	err := r.DB.NewUpdate().Model(&detail).Where("id = ?", data.Id).Scan(ctx)
	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int, userId int) error {
	_, err := r.DB.NewUpdate().
		Table("employees").
		Set("deleted_by = ?", userId).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).Exec(ctx)
	return err
}
