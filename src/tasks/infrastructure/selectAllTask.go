package infrastructure

import (
	"backend-golang-gin/tasks/domain"
	"log"

	"github.com/gin-gonic/gin"
)

func (ti *TaskInfrastructure) SelectAllTask(c *gin.Context) {
	var data []domain.Task

	rows, err := ti.DB.Query("SELECT * FROM tasks")

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var (
			id_task     string
			name_task   string
			status_task string
		)

		err := rows.Scan(&id_task, &name_task, &status_task)

		if err != nil {
			log.Println(err)
			return
		}

		task := domain.Task{ID: id_task, Name: name_task, Status: status_task}
		data = append(data, task)
	}

	selectAllResponseTask(data, c)
}
