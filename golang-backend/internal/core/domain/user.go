package domain

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID `gorm:"primaryKey"`
	UserName string    `gorm:"unique"`
	Email    string    `gorm:"unique"`
	Password string
	Tickets  []Ticket
}
