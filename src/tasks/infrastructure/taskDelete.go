package infrastructure

import (
	"backend-golang-gin/tasks/application"
	"log"

	"github.com/gin-gonic/gin"
)

func (ti *TaskInfrastructure) DeleteTask(c *gin.Context) {
	var ta application.TaskApplication

	id := ta.DeleteTask(c)

	stmt, err := ti.DB.Prepare("DELETE FROM tasks WHERE id_task=$1")

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(id)

	if err2 != nil {
		log.Println(err2)
	}

	deleteResponseTask(c)
}
