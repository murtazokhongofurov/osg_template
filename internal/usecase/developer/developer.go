package developer

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/developer"
)

type UseCase struct {
	developer Developer
}

func NewDevUseCase(developer Developer) *UseCase {
	return &UseCase{developer: developer}
}

func (du UseCase) CreateDeveloper(ctx context.Context, data developer.Create) (entity.Developer, error) {
	return du.developer.Create(ctx, data)
}

func (du UseCase) GetDeveloperById(ctx context.Context, id int) (developer.Detail, error) {
	data, err := du.developer.GetById(ctx, id)
	if err != nil {
		return developer.Detail{}, err
	}
	var detail developer.Detail
	detail.Id = data.Id
	detail.EmployeeId = data.EmployeeId
	detail.DeveloperRole = data.DeveloperRole
	return detail, nil
}

func (du UseCase) GetDeveloperList(ctx context.Context, filter developer.Filter) ([]developer.List, int, error) {
	data, count, err := du.developer.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	var list []developer.List
	for _, dev := range data {
		var detail developer.List
		detail.Id = dev.Id
		detail.EmployeeId = dev.EmployeeId
		detail.DeveloperRole = dev.DeveloperRole
		list = append(list, detail)
	}
	return list, count, nil
}

func (du UseCase) UpdateDeveloper(ctx context.Context, data developer.Update) (entity.Developer, error) {
	return du.developer.Update(ctx, data)
}

func (du UseCase) DeleteDeveloper(ctx context.Context, id int, userId int) error {
	return du.developer.Delete(ctx, id, userId)
}
