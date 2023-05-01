package comment

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/comment"
)

type UserCase struct {
	repo Comment
}

func NewCommentUseCase(repo Comment) *UserCase {
	return &UserCase{repo: repo}
}

func (cu UserCase) CreateComment(ctx context.Context, data comment.Create) (entity.Comment, error) {
	return cu.repo.Create(ctx, data)
}

func (cu UserCase) GetCommentList(ctx context.Context, filter comment.Filter) ([]entity.Comment, int, error) {
	data, count, err := cu.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	var list []entity.Comment
	for _, com := range data {
		var detail entity.Comment
		detail.Id = com.Id
		detail.DeveloperId = com.DeveloperId
		detail.Text = com.Text
		list = append(list, detail)
	}
	return list, count, nil
}

func (cu UserCase) GetTaskId(ctx context.Context, id int) ([]entity.Comment, int, error) {
	data, count, err := cu.repo.GetTaskId(ctx, id)
	if err != nil {
		return nil, 0, err
	}
	var list []entity.Comment
	for _, cmt := range data {
		var detail entity.Comment
		detail.Id = cmt.Id
		detail.DeveloperId = cmt.DeveloperId
		detail.TaskId = cmt.TaskId
		detail.Text = cmt.Text
		list = append(list, detail)
	}
	return list, count, nil
}

func (cu UserCase) UpdateComment(ctx context.Context, data comment.Update) (entity.Comment, error) {
	return cu.repo.Update(ctx, data)
}

func (cu UserCase) DeleteComment(ctx context.Context, id int, userId int) error {
	return cu.repo.Delete(ctx, id, userId)
}
