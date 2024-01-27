package application

import "github.com/gin-gonic/gin"

type TaskApplicationInterface interface {
	InsertTask(c *gin.Context)
}
