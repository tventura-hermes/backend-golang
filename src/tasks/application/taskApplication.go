package application

import (
	"database/sql"
)

type TaskApplication struct {
	DB *sql.DB
}

func NewTaskApplication(db *sql.DB) TaskApplicationInterface {
	return &TaskApplication{DB: db}
}
