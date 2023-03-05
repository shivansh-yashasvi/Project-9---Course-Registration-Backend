package structs

import "time"

type Data struct {
	Success bool `json:"success"`
	Data    Slot `json:"data"`
}

type Slot struct {
	ID    string    `json:"id"   validate:"required"`
	Day   string    `json:"day" validate:"required"`
	Start time.Time `json:"start" validate:"required"`
	End   time.Time `json:"end" validate:"required"`
}
