package infrastructure

import "github.com/gin-gonic/gin"

type App struct {
	Routes *gin.Engine
}

func (a *App) SendMessage() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	a.Routes = router
}

func (a *App) Run() {
	a.Routes.Run(":8080")
}
