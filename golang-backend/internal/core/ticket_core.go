package core

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/ticker"
	"github.com/google/uuid"
)

type TicketRepositoryIml struct {
	repo ports.TicketRepository
}

func NewTicketRepository(repo ports.TicketRepository) ports.TicketService {
	return &TicketRepositoryIml{repo: repo}
}

// Reserve seat and recieve ticket
func (r TicketRepositoryIml) ReqReserveSeat(input domain.InputTicket) (domain.Ticket, error) {
	newTicket := domain.Ticket{
		TicketId: uuid.New(),
		Status:   "Pending",
		UserId:   input.UserId,
		MovieId:  input.MovieId,
	}
	ticket, err := r.repo.CreateTicketAndReserveSeat(newTicket, input.TheaterId, input.SeatReserve)
	if err != nil {
		return ticket, err
	}
	return ticket, err
}

// Get ticket by id
func (r TicketRepositoryIml) ReqGetTicketById(id uuid.UUID) (domain.Ticket, error) {
	ticket, err := r.repo.ResTicketById(id)
	if err != nil {
		return ticket, err
	}
	return ticket, err
}
