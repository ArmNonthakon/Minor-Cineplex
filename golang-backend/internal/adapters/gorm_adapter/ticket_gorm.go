package adapters_gorm

import (
	"fmt"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/ticker"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewTicketGorm(db *gorm.DB) ports.TicketRepository {
	return &GormDB{db: db}
}

// Reserve seat and recieve ticket
func (g *GormDB) CreateTicketAndReserveSeat(newTicket domain.Ticket, UserName string, theaterId uuid.UUID, seatReserve []string) (domain.Ticket, error) {
	for _, seat := range seatReserve {
		if result := g.CheckSeat(seat, theaterId); !result {
			return newTicket, fmt.Errorf("CANT RESERVE")
		}

	}
	user := domain.User{}
	if result := g.db.Find(&user, "user_name = ?", UserName); result.Error != nil {
		return newTicket, result.Error
	}
	newTicket.UserId = user.UserId
	if result := g.db.Create(newTicket); result.Error != nil {
		return newTicket, result.Error
	}
	fmt.Print(theaterId)
	for _, seat := range seatReserve {
		g.ReserveSeat(seat, theaterId, newTicket.TicketId)

	}
	ticket, err := g.ResTicketById(newTicket.TicketId)
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}

// Get ticket by id
func (g *GormDB) ResTicketById(id uuid.UUID) (domain.Ticket, error) {
	ticket := domain.Ticket{}
	if result := g.db.Preload("Movie").Preload("Seats.Theater").Preload("User").Find(&ticket, "ticket_id = ?", id); result.Error != nil {
		return ticket, result.Error
	}
	return ticket, nil
}
