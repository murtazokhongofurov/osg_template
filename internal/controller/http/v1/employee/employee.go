package employee

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	employee_service "github.com/osg_template/internal/service/employee"
	employee "github.com/osg_template/internal/usecase/employee"
)

type EmployeeController struct {
	useCase *employee.UseCase
}

func NewEmployeeController(usecase *employee.UseCase) *EmployeeController {
	return &EmployeeController{useCase: usecase}
}

func (h EmployeeController) CreateEmployee(c *gin.Context) {
	var body employee_service.Create
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding request",
			"status":  false,
		})
		return
	}

	fmt.Println("id: ", body.CreatedBy)
	ctx := context.Background()
	data, err := h.useCase.CreateEmployee(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created employee",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
		},
	})
}

func (h EmployeeController) GetEmployeeById(c *gin.Context) {
	id := c.Param("id")
	empId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing id",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := h.useCase.GetDetail(ctx, empId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
		},
	})
}

func (h EmployeeController) GetAllEmployee(c *gin.Context) {
	var filter employee_service.Filter
	limitQ := c.Query("limit")
	if limitQ != "" {
		queryInt, err := strconv.Atoi(limitQ)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "limit must be number!",
				"status":  false,
			})
		}
		filter.Limit = &queryInt
	}
	offsetQ := c.Query("offset")
	if offsetQ != "" {
		queryInt, err := strconv.Atoi(offsetQ)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "offset must be number!",
				"status":  false,
			})
		}
		filter.Offset = &queryInt
	}

	ctx := context.Background()

	data, count, err := h.useCase.GetEmployeeList(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
			"count":  count,
		},
	})
}

func (h EmployeeController) UpdateEmployee(c *gin.Context) {
	var body employee_service.Update
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding request",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := h.useCase.UpdateEmployee(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
		},
	})
}

func (h EmployeeController) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	userId := c.Param("user_id")
	ctx := context.Background()

	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing id",
			"status":  false,
		})
		return
	}
	adminId, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing id",
			"status":  false,
		})
		return
	}
	err = h.useCase.DeleteEmployee(ctx, Id, adminId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
		"status":  true,
	})
}
