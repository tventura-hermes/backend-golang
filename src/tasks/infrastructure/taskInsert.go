package infrastructure

import (
	"backend-golang-gin/tasks/application"
	"log"

	"github.com/gin-gonic/gin"
)

func (m *TaskInfrastructure) InsertTask(c *gin.Context) {
	var ta application.TaskApplication

	post := ta.InsertTask(c)

	stmt, err := m.DB.Prepare("INSERT INTO tasks (id_task, name_task, status_task) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(post.ID, post.Name, post.Status)
	if err2 != nil {
		log.Println(err2)
	}

	insertResponseTask(&post, c)
}
