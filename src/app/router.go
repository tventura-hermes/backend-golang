package app

import (
	"backend-golang-gin/tasks/infrastructure"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type App struct {
	Routes *gin.Engine
	DB     *sql.DB
}

func (a *App) CreateRoutes() {
	router := gin.Default()

	it := infrastructure.NewTaskInfrastructure(a.DB)

	router.GET("/tasks", it.SelectAllTask)
	router.POST("/tasks", it.InsertTask)
	router.POST("/tasks/:id", it.UpdateTask)
	router.GET("/tasks/:id", it.DeleteTask)

	a.Routes = router
}

func (a *App) Run() {
	a.Routes.Run(":8080")
}
