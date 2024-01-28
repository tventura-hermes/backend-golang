package domain

import "github.com/gin-gonic/gin"

type TaskInterface interface {
	InsertTask(c *gin.Context)
	SelectAllTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}
