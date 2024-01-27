package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Routes *gin.Engine
}

func (a *App) CreateRoutes() {
	router := gin.Default()

	router.GET("/tasks", Connect)
	router.POST("/tasks", CreateTask)
	a.Routes = router
}

func (a *App) Run() {
	a.Routes.Run(":8080")
}
