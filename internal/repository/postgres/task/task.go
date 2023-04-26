package task

import (
	"context"
	"time"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/task"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data task.Create) (entity.Task, error) {
	now := time.Now()
	var detail entity.Task

	detail.Title = data.Title
	detail.Description = data.Description
	detail.DeveloperId = data.DeveloperId
	detail.FileUrl = data.FileUrl
	detail.CreatedAt = &now
	_, err := r.NewInsert().Model(&detail).Exec(ctx)
	return detail, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.Task, error) {
	var detail entity.Task

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	return detail, err
}

func (r Repository) GetAll(ctx context.Context, filter task.Filter) ([]entity.Task, int, error) {
	var list []entity.Task
	q := r.NewSelect().Model(&list)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}
	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}
	if filter.Search != nil {
		q.WhereGroup(" and ", func(query *bun.SelectQuery) *bun.SelectQuery {
			query.Where("text ilike ?", "%"+*filter.Search+"%")
			query.WhereOr("description ilike ?", "%"+*filter.Search+"%")
			return query
		})
	}

	count, err := q.ScanAndCount(ctx)
	return list, count, err
}

func (r Repository) Update(ctx context.Context, data task.Update) (entity.Task, error) {
	var detail entity.Task

	detail.Id = data.Id

	if data.Title != nil {
		detail.Title = data.Title
	}
	if data.Description != nil {
		detail.Description = data.Description
	}

	_, err := r.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int, userId int) error {
	_, err := r.NewUpdate().
		Table("tasks").
		Set("deleted_by = ?", userId).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
