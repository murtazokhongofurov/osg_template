package task

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	task_service "github.com/osg_template/internal/service/task"
	"github.com/osg_template/internal/usecase/task"
)

type Controller struct {
	userCase *task.UseCase
}

func NewConntroller(useCase *task.UseCase) *Controller {
	return &Controller{userCase: useCase}
}

func (tc Controller) CreateTask(c *gin.Context) {
	var body task_service.Create
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding request",
			"status":  false,
		})
		return
	}
	ctx := context.Background()

	data, err := tc.userCase.CreateTask(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
		},
	})
}

func (tc Controller) GetTaskList(c *gin.Context) {
	var filter task_service.Filter
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

	list, count, err := tc.userCase.GetTaskList(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  true,
		"data": map[string]interface{}{
			"result": list,
			"count":  count,
		},
	})

}

func (tc Controller) GetTaskId(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := tc.userCase.GetTaskDetail(ctx, taskId)
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
		},
	})
}

func (tc Controller) UpdateTask(c *gin.Context) {
	var body task_service.Update
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid info",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := tc.userCase.UpdateTask(ctx, body)
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

func (tc Controller) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid task id",
			"status":  false,
		})
		return
	}
	userId := c.Param("user_id")
	adminId, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	err = tc.userCase.DeleteTask(ctx, taskId, adminId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
		"status":  true,
	})
}
