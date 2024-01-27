package infrastructure

import (
	"backend-golang-gin/tasks/application"
	"backend-golang-gin/tasks/domain"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskInfrastructure struct {
	DB *sql.DB
}

func NewTaskInfrastructure(db *sql.DB) domain.TaskInterface {
	return &TaskInfrastructure{DB: db}
}

func (m *TaskInfrastructure) CreateTask(c *gin.Context) {
	var ta application.TaskApplication

	ta.InsertTask(c)

	response := m.InsertTask(application.TaskData)

	if response {
		c.JSON(http.StatusOK, gin.H{"message": "Dato insertado", "data": application.TaskData})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dato no insertado"})
	}
}
