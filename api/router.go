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
	r.GET("/employee/info/:id", option.Emp.GetEmployeeById)
	r.GET("/employees/:limit/:offset/:search", option.Emp.GetAllEmployee)
	r.PUT("/employee", option.Emp.UpdateEmployee)
	r.GET("/employee/:id", option.Emp.DeleteEmployee)

	// task
	r.POST("/task", option.Task.CreateTask)
	r.GET("/tasks", option.Task.GetTaskList)

	log.Fatalln(r.Run(":" + cfg.HttpPort))
}
