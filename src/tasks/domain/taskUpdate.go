package domain

type TaskUpdate struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
