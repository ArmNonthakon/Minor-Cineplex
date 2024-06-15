package domain

import "github.com/google/uuid"

type Ticket struct {
	TicketId uuid.UUID `gorm:"primaryKey"`
	Status   string    `gorm:"not null"`
	UserId   uuid.UUID `gorm:"size:191;index"`
	User     User
	MovieId  uuid.UUID `gorm:"size:191;index"`
	Movie    Movie
	Seats    []Seat
}
type InputTicket struct {
	UserId      uuid.UUID
	MovieId     uuid.UUID
	TheaterId   uuid.UUID
	SeatReserve []string
}
type InputTicketId struct {
	TicketId uuid.UUID
}
