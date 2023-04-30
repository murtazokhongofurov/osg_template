package project

import (
	"context"
	"time"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/project"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewProjectRepo(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data project.Create) (entity.Project, error) {
	var detail entity.Project
	now := time.Now()
	detail.EmployeeId = data.EmployeeId
	detail.Name = data.Name
	detail.StartedDate = data.StartedDate
	detail.FinishedDate = data.FinishedDate
	detail.Status = data.Status
	detail.FileUrl = data.FileUrl
	detail.CreatedAt = now

	_, err := r.DB.NewInsert().Model(&detail).Exec(ctx)
	return detail, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.Project, error) {
	var detail entity.Project
	err := r.DB.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)
	return detail, err
}

func (r Repository) GetAll(ctx context.Context, filter project.Filter) ([]entity.Project, int, error) {
	var list []entity.Project

	q := r.DB.NewSelect().Model(&list)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}
	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}
	if filter.Search != nil {
		q.WhereGroup(" and ", func(sq *bun.SelectQuery) *bun.SelectQuery {
			sq.Where("name ilike ?", "%"+*filter.Search+"%")
			return sq
		})
	}
	count, err := q.ScanAndCount(ctx)
	return list, count, err
}

func (r Repository) Update(ctx context.Context, data project.Update) (entity.Project, error) {
	var detail entity.Project
	detail.Id = data.Id
	if data.Name != nil {
		detail.Name = data.Name
	}
	if data.FileUrl != nil {
		detail.FileUrl = data.FileUrl
	}
	if data.Status != nil {
		detail.Status = data.Status
	}
	_, err := r.DB.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int, userId int) error {
	_, err := r.DB.NewUpdate().
		Table("projects").
		Set("deleted_by = ?", userId).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).Exec(ctx)

	return err
}
