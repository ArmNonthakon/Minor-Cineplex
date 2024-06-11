package domain

import "time"

type Theater struct {
	TheaterId string `gorm:"primaryKey"`
	StartAt   time.Time
	EndAt     time.Time
	IsFull    bool
	Status    string
	MovieId   string
	Movie     Movie `gorm:"references:MovieID"`
}
