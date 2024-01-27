package application

import (
	"backend-golang-gin/tasks/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *TaskApplication) InsertTask(c *gin.Context) domain.Task {
	var post domain.Task

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return post
	}

	fmt.Println(post)
	return post
}
