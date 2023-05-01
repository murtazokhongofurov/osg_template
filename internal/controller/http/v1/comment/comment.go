package comment

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	com "github.com/osg_template/internal/service/comment"
	"github.com/osg_template/internal/usecase/comment"
)

type ComController struct {
	com *comment.UserCase
}

func NewCommentController(com *comment.UserCase) *ComController {
	return &ComController{com: com}
}

func (h ComController) CreateComment(c *gin.Context) {
	var body com.Create
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid info",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	ctx := context.Background()
	data, err := h.com.CreateComment(ctx, body)
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

func (h ComController) GetCommentTaskId(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing id",
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	data, count, err := h.com.GetTaskId(ctx, taskId)
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
			"count":  count,
		},
	})
}

func (h ComController) GetCommentList(c *gin.Context) {
	var filter com.Filter
	limit := c.Query("limit")
	if limit != "" {
		queryInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "parsing limit",
				"status":  false,
			})
			return
		}
		filter.Limit = &queryInt
	}
	offset := c.Query("offset")
	if offset != "" {
		queryInt, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "parsing limit",
				"status":  false,
			})
			return
		}
		filter.Limit = &queryInt
	}
	ctx := context.Background()
	data, count, err := h.com.GetCommentList(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
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
			"count":  count,
		},
	})
}

func (h ComController) UpdateComment(c *gin.Context) {
	var body com.Update
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid info",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	ctx := context.Background()
	fmt.Println(">>>>>>>>>>> ", body.Id, body.Text)
	data, err := h.com.UpdateComment(ctx, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
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

func (h ComController) DeleteComment(c *gin.Context) {
	comId := c.Param("id")
	adId := c.Param("user_id")
	id, err := strconv.Atoi(comId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing id",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	userId, err := strconv.Atoi(adId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing id",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	ctx := context.Background()
	err = h.com.DeleteComment(ctx, id, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"status":  false,
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
		"status":  true,
	})
}
