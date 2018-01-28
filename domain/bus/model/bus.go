package model

import (
	"time"

	"github.com/rodrigo-brito/bus-api-go/domain/schedule/model"
)

type Bus struct {
	ID         int64             `json:"id"`
	Number     *string           `json:"number"`
	Name       string            `json:"name"`
	Fare       float64           `json:"fare"`
	LastUpdate *time.Time        `json:"last_update"`
	Schedules  []*model.Schedule `json:"schedules,omitempty"`
}

func (b *Bus) IsEmpty() bool {
	return b == nil || (b.ID == 0 && b.Number == nil &&
		b.Name == "" && len(b.Schedules) == 0)
}
