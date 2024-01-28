package infrastructure

import (
	"backend-golang-gin/tasks/application"
	"log"

	"github.com/gin-gonic/gin"
)

func (ti *TaskInfrastructure) UpdateTask(c *gin.Context) {
	var ta application.TaskApplication

	post := ta.UpdateTask(c)

	stmt, err := ti.DB.Prepare("UPDATE tasks SET name_task=$1, status_task=$2 WHERE id_task=$3")

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(post.Name, post.Status, application.TaskUpdateId)

	if err2 != nil {
		log.Println(err2)
	}

	updateResponseTask(&post, c)
}
