package models

import (
	"time"
)

type Tracker struct {
	Timestamp time.Time    	`gorm:"type:timestamp" json:"timestamp"`
	VehicleId int    		`gorm:"type:int" json:"vehicleId"`
	Latitude  float64 		`gorm:"type:double" json:"latitude"`
	Longitude float64 		`gorm:"type:double" json:"longitude"`
}
