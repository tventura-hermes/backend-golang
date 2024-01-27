package infrastructure

import (
	"backend-golang-gin/tasks/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	application.CreateTask(c)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Insercion", "data": application.GlobalData})
}
