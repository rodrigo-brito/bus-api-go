package model

import "github.com/rodrigo-brito/bus-api-go/domain/bus/model"

type Company struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	ImageURL    string       `json:"imageUrl"`
	Description string       `json:"description"`
	Bus         []*model.Bus `json:"bus,omitempty"`
}
