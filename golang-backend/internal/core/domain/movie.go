package domain

import "github.com/google/uuid"

type Movie struct {
	MovieId  uuid.UUID `gorm:"primaryKey"`
	Title    string    `gorm:"unique"`
	Duration float32
	Theaters []Theater
	Tickets  []Ticket
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
	Title    string
	Duration float32
}
type InputTitle struct {
	Title string
}
