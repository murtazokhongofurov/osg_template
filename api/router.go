package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	e "github.com/osg_template/internal/controller/http/v1/employee"
	t "github.com/osg_template/internal/controller/http/v1/task"
	"github.com/osg_template/internal/pkg/config"
)

func Router(task *t.Controller, emp *e.EmployeeController) {
	r := gin.Default()

	cfg := config.Load()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})
	// employee
	r.POST("/employee", emp.CreateEmployee)
	r.GET("/employee/info/:id", emp.GetEmployeeById)
	r.GET("/employees/:limit/:offset/:search", emp.GetAllEmployee)
	r.PUT("/employee", emp.UpdateEmployee)
	r.GET("/employee/:id", emp.DeleteEmployee)

	// task
	r.POST("/task", task.CreateTask)
	r.GET("/tasks", task.GetTaskList)

	log.Fatalln(r.Run(":" + cfg.HttpPort))
}
