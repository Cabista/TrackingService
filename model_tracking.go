package main

import "time"

type Tracking struct {
	ID uint64 `json:"id,omitempty" gorm:"primary_key"`

	DriverID uint64 `json:"driver_id,omitempty"`

	LonCoord float64

	LatCoord float64

	CreatedAt time.Time

	DeletedAt *time.Time `sql:"index"`
}
