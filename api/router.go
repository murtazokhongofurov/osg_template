package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	d "github.com/osg_template/internal/controller/http/v1/developer"
	e "github.com/osg_template/internal/controller/http/v1/employee"
	p "github.com/osg_template/internal/controller/http/v1/project"
	t "github.com/osg_template/internal/controller/http/v1/task"
	"github.com/osg_template/internal/pkg/config"
)

type Options struct {
	Cfg  config.Config
	Emp  e.EmployeeController
	Task t.Controller
	Dev  d.DevController
	Pro  p.ProjectController
}

func Router(option *Options) {
	r := gin.Default()

	cfg := config.Load()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})
	// employee
	r.POST("/employee", option.Emp.CreateEmployee)
	r.GET("/employee/:id", option.Emp.GetEmployeeById)
	r.GET("/employees", option.Emp.GetAllEmployee)
	r.PUT("/employee", option.Emp.UpdateEmployee)
	r.GET("/employee/:id/:user_id", option.Emp.DeleteEmployee)

	// developer
	r.POST("/developer", option.Dev.CreateDev)
	r.GET("/developer/:id", option.Dev.GetDevId)
	r.GET("/developers", option.Dev.GetDevList)
	r.PUT("/developer", option.Dev.Update)
	r.DELETE("/developer/:id/:user_id", option.Dev.DeleteDeveloper)

	// project
	r.POST("/project", option.Pro.CreateProject)
	r.GET("/project/:id", option.Pro.GetAllProject)
	r.GET("/projects", option.Pro.GetAllProject)
	r.PUT("project", option.Pro.UpdateProject)
	r.DELETE("/project/:id/:user_id", option.Pro.DeleteProject)

	// task
	r.POST("/task", option.Task.CreateTask)
	r.GET("/tasks", option.Task.GetTaskList)
	r.GET("task/:id", option.Task.GetTaskId)
	r.PUT("task", option.Task.UpdateTask)
	r.DELETE("task/:id/:user_id", option.Task.DeleteTask)

	log.Fatalln(r.Run(":" + cfg.HttpPort))
}
