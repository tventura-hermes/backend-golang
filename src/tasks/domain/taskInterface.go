package domain

import "github.com/gin-gonic/gin"

type TaskInterface interface {
	InsertTask(post Task) bool
	CreateTask(c *gin.Context)
}
