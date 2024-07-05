package ports

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

type TicketService interface {
	ReqReserveSeat(input domain.InputTicket) (domain.Ticket, error)
	ReqGetTicketById(id uuid.UUID) (domain.Ticket, error)
	ReqGetTicketByUserName(userName string) ([]domain.Ticket, error)
}
