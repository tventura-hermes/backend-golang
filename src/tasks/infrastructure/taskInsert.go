package infrastructure

import (
	"backend-golang-gin/tasks/domain"
	"log"
)

func (m *TaskInfrastructure) InsertTask(post domain.Task) bool {
	log.Println(post.ID + " " + post.Name + " " + post.Status)

	stmt, err := m.DB.Prepare("INSERT INTO tasks (id_task, name_task, status_task) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println(err)
		return false
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(post.ID, post.Name, post.Status)
	if err2 != nil {
		log.Println(err2)
		return false
	}

	return true
}
