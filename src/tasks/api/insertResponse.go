package api

import (
	"time"
)

type ApiConfirmation struct {
	Data     ConfirmationData     `json:"data"`
	Included ConfirmationIncluded `json:"included"`
}

type ConfirmationResponse struct {
	Data ConfirmationData `json:"data"`
}

type ConfirmationData struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes ConfirmationAttributes `json:"attributes"`
	Relations  ConfirmationRelations  `json:"relations"`
}

type ConfirmationAttributes struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamps"`
}

type ConfirmationRelations struct {
	Included Included `json:"included"`
}

type Included struct {
	Data DataIncluded `json:"data"`
}

type DataIncluded struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type ConfirmationIncluded struct {
	Type       string                         `json:"type"`
	ID         string                         `json:"id"`
	Attributes ConfirmationIncludedAttributes `json:"attributes"`
}

type ConfirmationIncludedAttributes struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
