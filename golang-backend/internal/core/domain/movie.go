package domain

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	MovieId     uuid.UUID `gorm:"primaryKey"`
	Title       string    `gorm:"unique"`
	Genre       string
	PosterUrl   string
	ReleaseDate time.Time
	Duration    int
	Theaters    []Theater
	Tickets     []Ticket
}

type OnlyMovie struct {
	MovieId uuid.UUID `gorm:"primaryKey"`
	Title   string
}

type MovieWithTheater struct {
	MovieId  uuid.UUID `gorm:"primaryKey"`
	Title    string
	Theaters []Theater
}

type InputMovie struct {
	Title       string
	PosterUrl   string
	Duration    int
	Genre       string
	ReleaseDate time.Time
}
type InputTitle struct {
	Title string
}
