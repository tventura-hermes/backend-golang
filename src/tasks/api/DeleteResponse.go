package api

type DeleteConfirmationResponse struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes ConfirmationAttributes `json:"attributes"`
}
