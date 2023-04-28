package main

import (
	"github.com/osg_template/api"
	dev_controller "github.com/osg_template/internal/controller/http/v1/developer"
	emp_controller "github.com/osg_template/internal/controller/http/v1/employee"
	pro_controller "github.com/osg_template/internal/controller/http/v1/project"
	task_controller "github.com/osg_template/internal/controller/http/v1/task"
	"github.com/osg_template/internal/pkg/config"
	"github.com/osg_template/internal/pkg/repository"
	"github.com/osg_template/internal/repository/postgres/developer"
	"github.com/osg_template/internal/repository/postgres/employee"
	"github.com/osg_template/internal/repository/postgres/project"
	"github.com/osg_template/internal/repository/postgres/task"
	developer_service "github.com/osg_template/internal/service/developer"
	employee_service "github.com/osg_template/internal/service/employee"
	project_service "github.com/osg_template/internal/service/project"
	task_service "github.com/osg_template/internal/service/task"
	dev_use_case "github.com/osg_template/internal/usecase/developer"
	emp_use_case "github.com/osg_template/internal/usecase/employee"
	pro_use_case "github.com/osg_template/internal/usecase/project"
	task_use_case "github.com/osg_template/internal/usecase/task"
)

func main() {

	cfg := config.Load()
	// database
	postgresDB := repository.NewPostgres()

	// repository
	employeeRepo := employee.NewRepository(postgresDB)
	taskRepo := task.NewRepository(postgresDB)
	projectRepo := project.NewProjectRepo(postgresDB)
	developerRepo := developer.NewDeveloperRepo(postgresDB)

	// service
	taskService := task_service.NewService(taskRepo)
	employeeService := employee_service.NewService(employeeRepo)
	projectService := project_service.NewProjectService(projectRepo)
	developerService := developer_service.NewDeveloperService(developerRepo)

	// use case
	taskUseCase := task_use_case.NewUseCase(taskService)
	employeeUseCase := emp_use_case.NewUseCase(employeeService)
	projectUseCase := pro_use_case.NewProjectUseCase(projectService)
	developerUseCase := dev_use_case.NewDevUseCase(developerService)

	// controller
	taskController := task_controller.NewConntroller(taskUseCase)
	empController := emp_controller.NewEmployeeController(employeeUseCase)
	projectContr := pro_controller.NewProjectController(projectUseCase)
	devController := dev_controller.NewDevController(developerUseCase)

	api.Router(&api.Options{
		Cfg:  cfg,
		Emp:  *empController,
		Task: *taskController,
		Dev:  *devController,
		Pro:  *projectContr,
	})

}
