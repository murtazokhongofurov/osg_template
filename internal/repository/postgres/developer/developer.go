package developer

import (
	"context"
	"time"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/developer"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewDeveloperRepo(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data developer.Create) (entity.Developer, error) {
	var detail entity.Developer
	detail.EmployeeId = data.EmployeeId
	detail.DeveloperRole = data.DeveloperRole
	now := time.Now()
	detail.CreatedAt = &now
	_, err := r.DB.NewInsert().Model(&detail).Exec(ctx)
	return detail, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.Developer, error) {
	var detail entity.Developer
	err := r.DB.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)
	return detail, err
}

func (r Repository) GetAll(ctx context.Context, filter developer.Filter) ([]entity.Developer, int, error) {
	var list []entity.Developer
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

func (r Repository) Update(ctx context.Context, data developer.Update) (entity.Developer, error) {
	var detail entity.Developer
	detail.Id = data.Id

	if data.DeveloperRole != nil {
		detail.DeveloperRole = data.DeveloperRole
	}
	_, err := r.DB.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)
	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int, userId int) error {
	_, err := r.DB.NewUpdate().
		Table("developers").
		Set("deleted_by = ?", userId).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).Exec(ctx)
	return err
}
