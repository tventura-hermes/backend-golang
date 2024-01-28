package application

import (
	"backend-golang-gin/tasks/domain"

	"github.com/gin-gonic/gin"
)

var TaskUpdateId string

func (*TaskApplication) UpdateTask(c *gin.Context) domain.TaskUpdate {
	var post domain.TaskUpdate

	if err := c.BindJSON(&post); err != nil {
		return post
	}

	//validaciones de datos

	TaskUpdateId = c.Param("id")

	return post
}
