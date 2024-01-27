package application

import (
	"backend-golang-gin/tasks/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var GlobalData domain.Task

func CreateTask(c *gin.Context) {
	var newTask domain.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(newTask)

	GlobalData = newTask
}
