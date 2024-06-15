package domain

import "github.com/google/uuid"

type Seat struct {
	SeatID     uuid.UUID `gorm:"primaryKey"`
	SeatNumber string
	TheaterId  uuid.UUID `gorm:"size:191;index"`
	Theater    Theater
	TicketId   uuid.UUID `gorm:"size:191;index"`
	Ticket     Ticket
}
