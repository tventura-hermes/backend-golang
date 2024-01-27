package infrastructure

import (
	"backend-golang-gin/tasks/api"
	"backend-golang-gin/tasks/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *TaskInfrastructure) ResponseTask(c *gin.Context) {

	var ta application.TaskApplication

	TaskData := ta.InsertTask(c)

	response := m.InsertTask(TaskData)

	jsonapi := api.ApiConfirmation{
		Data: api.ConfirmationResponse{
			Data: api.ConfirmationData{
				Type: "confirmation",
				ID:   "1",
				Attributes: api.ConfirmationAttributes{
					Message:   "La inserción a la base de datos se ha realizado con éxito.",
					Timestamp: time.Now(),
				},
				Relations: api.ConfirmationRelations{
					Included: api.Included{
						Data: api.DataIncluded{
							Type: "Task",
							ID:   TaskData.ID,
						},
					},
				},
			},
		},
		Included: api.ConfirmationIncluded{
			Type: "Task",
			ID:   TaskData.ID,
			Attributes: api.ConfirmationIncludedAttributes{
				Name:   TaskData.Name,
				Status: TaskData.Status,
			},
		},
	}

	if response {
		c.JSON(http.StatusCreated, jsonapi)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dato no insertado"})
	}
}
