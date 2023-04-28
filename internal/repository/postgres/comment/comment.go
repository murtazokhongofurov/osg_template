package comment

import (
	"context"
	"time"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/comment"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewCommentRepo(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data comment.Create) (entity.Comment, error) {
	var detail entity.Comment
	now := time.Now()
	detail.DeveloperId = data.DeveloperId
	detail.Text = data.Text
	detail.CreatedAt = &now
	_, err := r.DB.NewInsert().Model(&detail).Exec(ctx)
	return detail, err
}

func (r Repository) GetAll(ctx context.Context, filter comment.Filter) ([]entity.Comment, int, error) {
	var list []entity.Comment
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

func (r Repository) GetTaskId(ctx context.Context, id int) ([]entity.Comment, int, error) {
	var list []entity.Comment
	q := r.DB.NewSelect().Model(&list).Where("task_id = ?", id)
	count, err := q.ScanAndCount(ctx)
	return list, count, err
}

func (r Repository) Update(ctx context.Context, data comment.Update) (entity.Comment, error) {
	var detail entity.Comment
	detail.Id = data.Id
	if data.Text != nil {
		detail.Text = data.Text
	}
	_, err := r.DB.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)
	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int, userId int) error {
	_, err := r.DB.NewUpdate().
		Table("comment").
		Set("deleted_by = ?", userId).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", id).Exec(ctx)
	return err
}
