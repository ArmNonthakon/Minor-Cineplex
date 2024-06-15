package domain

import (
	"time"

	"github.com/google/uuid"
)

type Theater struct {
	TheaterId     uuid.UUID `gorm:"primaryKey"`
	TheaterNumber string
	MovieId       uuid.UUID `gorm:"size:191;index"`
	StartTime     time.Time
	SeatRow       int
	SeatCol       int
	Movie         Movie
	Seats         []Seat
}
type InputTheater struct {
	TheaterNumber string
	MovieId       uuid.UUID
	StartTime     time.Time
	SeatRow       int
	SeatCol       int
}
type InputNumber struct {
	TheaterNumber string
}
type InputTheaterId struct {
	TheaterId string
}
