package model

import "github.com/rodrigo-brito/bus-api-go/domain/schedule/model"

type Bus struct {
	ID        int64             `json:"id"`
	Number    *int64            `json:"number"`
	Name      string            `json:"name"`
	Fare      float64           `json:"fare"`
	Schedules []*model.Schedule `json:"schedules,omitempty"`
}
