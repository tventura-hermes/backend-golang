package application

import (
	"backend-golang-gin/tasks/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var TaskData domain.Task

func (t *TaskApplication) InsertTask(c *gin.Context) {
	var post domain.Task

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(post)

	TaskData = post
}
