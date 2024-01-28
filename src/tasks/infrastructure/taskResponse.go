package infrastructure

import (
	"backend-golang-gin/tasks/api"
	"backend-golang-gin/tasks/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func insertResponseTask(TaskData *domain.Task, c *gin.Context) {

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

	c.JSON(http.StatusCreated, jsonapi)
}

func selectAllResponseTask(data []domain.Task, c *gin.Context) {
	var jsonapi []api.SelectResponse

	for _, a := range data {
		format := api.SelectResponse{
			Type: "task",
			ID:   a.ID,
			Attributes: api.ConfirmationIncludedAttributes{
				Name:   a.Name,
				Status: a.Status,
			},
		}
		jsonapi = append(jsonapi, format)
	}

	c.JSON(http.StatusFound, gin.H{"data": jsonapi})
}
