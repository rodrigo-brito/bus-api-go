package model

import "github.com/rodrigo-brito/bus-api-go/domain/bus/model"

type Company struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	ImageURL    string       `json:"imageUrl"`
	Description string       `json:"description"`
	Bus         []*model.Bus `json:"bus,omitempty"`
}

func (c *Company) IsEmpty() bool {
	return c == nil || (c.ID == 0 && c.Name == "" && c.ImageURL == "" &&
		c.Description == "" && len(c.Bus) == 0)
}
