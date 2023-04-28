package project

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/project"
)

type UseCase struct {
	project Prject
}

func NewProjectUseCase(project Prject) *UseCase {
	return &UseCase{project: project}
}

func (pu UseCase) CreateProject(ctx context.Context, data project.Create) (entity.Project, error) {
	return pu.project.Create(ctx, data)
}

func (pu UseCase) GetProjectById(ctx context.Context, id int) (project.Detail, error) {
	data, err := pu.project.GetById(ctx, id)
	if err != nil {
		return project.Detail{}, err
	}
	var detail project.Detail
	detail.Id = &data.Id
	detail.Name = data.Name
	detail.FileUrl = data.FileUrl
	detail.StartedDate = data.StartedDate
	detail.FinishedDate = data.FinishedDate
	detail.Status = data.Status
	return detail, nil
}

func (pu UseCase) UpdateProject(ctx context.Context, data project.Update) (entity.Project, error) {
	return pu.project.Update(ctx, data)
}

func (pu UseCase) GetProjectList(ctx context.Context, filter project.Filter) ([]project.List, int, error) {
	data, count, err := pu.project.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	var list []project.List
	for _, pro := range data {
		var detail project.List
		detail.Id = &pro.Id
		detail.Name = pro.Name
		detail.StartedDate = pro.StartedDate
		detail.FinishedDate = pro.FinishedDate
		detail.Status = pro.Status
		list = append(list, detail)
	}
	return list, count, nil
}

func (pu UseCase) DeleteProject(ctx context.Context, id int, userId int) error {
	return pu.project.Delete(ctx, id, userId)
}
