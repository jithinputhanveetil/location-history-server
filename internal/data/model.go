package data

import "time"

type History struct {
	Lat           float32    `json:"lat,omitempty"`
	Lng           float32    `json:"lng,omitempty"`
	InsertionTime *time.Time `json:"-"`
}

type Location struct {
	OrderID string
	History []History
}
