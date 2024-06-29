package adapters_gorm

import (
	"fmt"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

func (g *GormDB) ReserveSeat(number string, theaterId uuid.UUID, ticketId uuid.UUID) bool {
	seat := domain.Seat{
		SeatID:     uuid.New(),
		SeatNumber: number,
		TheaterId:  theaterId,
		TicketId:   ticketId,
	}
	if result := g.db.Create(seat); result.Error != nil {
		fmt.Print("g.db.Create()")
		return false
	}
	return true
}
func (g *GormDB) CheckSeat(number string, theaterId uuid.UUID) bool {
	seat := &domain.Seat{}
	if result := g.db.First(&seat, "theater_id = ? and seat_number = ?", theaterId, number); result.Error != nil {
		return true
	}
	return false
}
