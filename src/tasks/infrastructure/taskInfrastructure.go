package infrastructure

import (
	"backend-golang-gin/tasks/domain"
	"database/sql"
)

type TaskInfrastructure struct {
	DB *sql.DB
}

func NewTaskInfrastructure(db *sql.DB) domain.TaskInterface {
	return &TaskInfrastructure{DB: db}
}
