package developer

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dev "github.com/osg_template/internal/service/developer"
	"github.com/osg_template/internal/usecase/developer"
)

type DevController struct {
	dev *developer.UseCase
}

func NewDevController(dev *developer.UseCase) *DevController {
	return &DevController{dev: dev}
}

func (h DevController) CreateDev(c *gin.Context) {
	var body dev.Create
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding request",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := h.dev.CreateDeveloper(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
		},
	})
}

func (h DevController) GetDevId(c *gin.Context) {
	id := c.Param("id")
	devId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parsing error",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := h.dev.GetDeveloperById(ctx, devId)
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

func (h DevController) Update(c *gin.Context) {
	var body dev.Update
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding request",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, err := h.dev.UpdateDeveloper(ctx, body)
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

func (h DevController) GetDevList(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	var filter dev.Filter
	if limit != "" {
		limitI, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "parsing limit",
				"status":  false,
			})
		}
		filter.Limit = &limitI
	}
	if offset != "" {
		offsetI, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "parsing offset",
				"status":  false,
			})
		}
		filter.Offset = &offsetI
	}
	ctx := context.Background()
	data, count, err := h.dev.GetDeveloperList(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Ok",
		"status":  true,
		"data": map[string]interface{}{
			"result": data,
			"count":  count,
		},
	})
}

func (h DevController) DeleteDeveloper(c *gin.Context) {
	id := c.Param("id")
	useId := c.Param("user_id")
	devId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parsing id",
			"status":  false,
		})
		return
	}
	userId, err := strconv.Atoi(useId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parsing user_id",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	err = h.dev.DeleteDeveloper(ctx, devId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully delete",
		"status":  true,
	})
}
