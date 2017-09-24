package model

type Schedule struct {
	ID          int     `json:"id"`
	Origin      string  `json:"origin"`
	Destiny     string  `json:"destiny"`
	Observation *string `json:"observation"`
	Time        string  `json:"time"`
}
