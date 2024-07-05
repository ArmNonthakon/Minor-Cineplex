package ports

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

type TicketRepository interface {
	CreateTicketAndReserveSeat(newTicket domain.Ticket, userName string, theaterId uuid.UUID, seatReserve []string) (domain.Ticket, error)
	ResTicketById(id uuid.UUID) (domain.Ticket, error)
	ResTicketByUserName(userName string) ([]domain.Ticket, error)
}
