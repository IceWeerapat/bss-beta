package home

import "time"

type Status struct {
	Time time.Time `json:"time"`
}

func NewStatus() *Status {
	return &Status{time.Now()}
}
