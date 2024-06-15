package adapters_gorm

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

func (g *GormDB) ReserveSeat(number string, theaterId uuid.UUID, ticketId uuid.UUID) error {
	seat := domain.Seat{
		SeatID:     uuid.New(),
		SeatNumber: number,
		TheaterId:  theaterId,
		TicketId:   ticketId,
	}
	if result := g.db.Create(seat); result.Error != nil {
		return result.Error
	}
	return nil
}
