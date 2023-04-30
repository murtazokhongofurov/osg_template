package project

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pro "github.com/osg_template/internal/service/project"
	"github.com/osg_template/internal/usecase/project"
)

type ProjectController struct {
	pro *project.UseCase
}

func NewProjectController(pro *project.UseCase) *ProjectController {
	return &ProjectController{pro: pro}
}

func (h ProjectController) CreateProject(c *gin.Context) {
	var body pro.Create
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := h.pro.CreateProject(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
		},
	})
}

func (h *ProjectController) GetProject(c *gin.Context) {
	id := c.Param("id")
	proId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parsing id",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	info, err := h.pro.GetProjectById(ctx, proId)
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
			"result": info,
		},
	})
}

func (h ProjectController) GetAllProject(c *gin.Context) {
	var filter pro.Filter
	limit := c.Query("limit")
	offset := c.Query("offset")
	if limit != "" {
		queryInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "parsing error",
				"status":  false,
			})
		}
		filter.Limit = &queryInt
	}
	if offset != "" {
		queryInt, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "parsing error",
				"status":  false,
			})
			return
		}
		filter.Offset = &queryInt
	}
	ctx := context.Background()
	list, count, err := h.pro.GetProjectList(ctx, filter)
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
			"result": list,
			"count":  count,
		},
	})
}

func (h ProjectController) UpdateProject(c *gin.Context) {
	var body pro.Update
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "binding request",
			"status":  false,
		})
		return
	}

	ctx := context.Background()
	data, err := h.pro.UpdateProject(ctx, body)
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

func (h ProjectController) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	userId := c.Param("user_id")
	proId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
			"status":  false,
		})
		return
	}
	adminId, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user_id",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	err = h.pro.DeleteProject(ctx, proId, adminId)
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
