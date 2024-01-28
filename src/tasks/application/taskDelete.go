package application

import "github.com/gin-gonic/gin"

func (*TaskApplication) DeleteTask(c *gin.Context) string {
	id := c.Param("id")

	return id
}
