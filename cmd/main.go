package main

import (
	emp_controller "github.com/osg_template/internal/controller/http/v1/employee"
	task_controller "github.com/osg_template/internal/controller/http/v1/task"
	"github.com/osg_template/internal/pkg/repository"
	"github.com/osg_template/internal/repository/postgres/employee"
	"github.com/osg_template/internal/repository/postgres/task"
	employee_service "github.com/osg_template/internal/service/employee"
	task_service "github.com/osg_template/internal/service/task"

	"github.com/osg_template/api"
	emp_use_case "github.com/osg_template/internal/usecase/employee"
	task_use_case "github.com/osg_template/internal/usecase/task"
)

func main() {

	// database
	postgresDB := repository.NewPostgres()

	// repository
	employeeRepo := employee.NewRepository(postgresDB)
	taskRepo := task.NewRepository(postgresDB)

	// service
	taskService := task_service.NewService(taskRepo)
	employeeService := employee_service.NewService(employeeRepo)

	// use case
	taskUseCase := task_use_case.NewUseCase(taskService)
	employeeUseCase := emp_use_case.NewUseCase(employeeService)

	taskController := task_controller.NewConntroller(taskUseCase)
	empController := emp_controller.NewEmployeeController(employeeUseCase)

	api.Router(taskController, empController)

}
