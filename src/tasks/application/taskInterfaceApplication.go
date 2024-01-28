package application

import (
	"backend-golang-gin/tasks/domain"

	"github.com/gin-gonic/gin"
)

type TaskApplicationInterface interface {
	InsertTask(c *gin.Context) domain.Task
	UpdateTask(c *gin.Context) domain.TaskUpdate
	DeleteTask(c *gin.Context) string
}
