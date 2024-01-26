package application

import "backend-golang-gin/tasks/domain"

var taskList = []domain.Task{
	{ID: "65b41bc5fc13ae2b9eb4c76e", Name: "Tarea importante", Status: "Doing"},
}

func AddTask(ir domain.TaskInterface) {
	ir.DataTask(taskList)
}
