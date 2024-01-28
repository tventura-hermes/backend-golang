package api

type SelectResponse struct {
	Type       string                         `json:"type"`
	ID         string                         `json:"id"`
	Attributes ConfirmationIncludedAttributes `json:"attributes"`
}
