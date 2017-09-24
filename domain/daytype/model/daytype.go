package model

import "github.com/rodrigo-brito/bus-api-go/domain/schedule/model"

type DayType struct {
	ID        int64             `json:"id"`
	Name      string            `json:"name"`
	Active    bool              `json:"active"`
	Schedules []*model.Schedule `json:"schedules,omitempty"`
}
