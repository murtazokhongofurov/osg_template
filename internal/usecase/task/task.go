package task

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/task"
)

type UseCase struct {
	task Task
}

func NewUseCase(task Task) *UseCase {
	return &UseCase{task: task}
}

func (tu UseCase) CreateTask(ctx context.Context, data task.Create) (entity.Task, error) {
	return tu.task.Create(ctx, data)
}

func (tu UseCase) UpdateTask(ctx context.Context, data task.Update) (entity.Task, error) {
	return tu.task.Update(ctx, data)
}

func (tu UseCase) GetTaskList(ctx context.Context, filter task.Filter) ([]task.List, int, error) {
	data, count, err := tu.task.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []task.List

	for _, t := range data {
		var detail task.List
		detail.Id = t.Id
		detail.Title = t.Title
		detail.Description = t.Description
		detail.FileUrl = t.FileUrl
		list = append(list, detail)
	}
	return list, count, nil

}

func (tu UseCase) GetTaskDetail(ctx context.Context, id int) (task.Detail, error) {
	data, err := tu.task.GetById(ctx, id)
	if err != nil {
		return task.Detail{}, err
	}

	var detail task.Detail
	detail.Id = data.Id
	detail.Title = data.Title
	detail.Description = data.Description
	detail.FileUrl = data.FileUrl

	return detail, nil
}

func (tu UseCase) DeleteTask(ctx context.Context, id int, userId int) error {
	return tu.task.Delete(ctx, id, userId)
}
