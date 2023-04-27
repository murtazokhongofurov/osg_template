package usecase

import (
	"context"

	"github.com/osg_template/internal/entity"
	"github.com/osg_template/internal/service/employee"
)

type UseCase struct {
	employee Employee
}

func NewUseCase(employee Employee) *UseCase {
	return &UseCase{employee: employee}
}

func (e UseCase) CreateEmployee(ctx context.Context, data employee.Create) (entity.Employee, error) {

	return e.employee.Create(ctx, data)
}

func (e UseCase) GetDetail(ctx context.Context, id int) (employee.Detail, error) {
	data, err := e.employee.GetById(ctx, id)
	if err != nil {
		return employee.Detail{}, err
	}
	var detail employee.Detail
	detail.FullName = data.FullName
	detail.ProfilePhoto = data.ProfilePhoto
	detail.BirthDate = data.BirthDate
	detail.Phone = data.Phone
	detail.Position = data.Position
	detail.Role = data.Role

	return detail, nil

}

func (e UseCase) GetEmployeeList(ctx context.Context, filter employee.Filter) ([]employee.List, int, error) {
	data, count, err := e.employee.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	var list []employee.List

	for _, em := range data {
		var detail employee.List
		detail.Id = &em.Id
		detail.FullName = em.FullName
		detail.ProfilePhoto = em.ProfilePhoto
		detail.Position = em.Position
		detail.Role = em.Role
		list = append(list, detail)
	}

	return list, count, nil
}

func (e UseCase) DeleteEmployee(ctx context.Context, id int, userId int) error {
	return e.employee.Delete(ctx, id, userId)
}

func (e UseCase) UpdateEmployee(ctx context.Context, data employee.Update) (entity.Employee, error) {
	return e.employee.Update(ctx, data)
}
