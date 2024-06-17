package adapters_gorm

import (
	"fmt"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

func (g *GormDB) ReserveSeat(number string, theaterId uuid.UUID, ticketId uuid.UUID) error {
	checkSeat := []domain.Seat{}
	if result := g.db.First(&checkSeat, "theater_id = ? and number = ?", theaterId, number); result.Error != nil {
		return result.Error
	}
	if len(checkSeat) != 0 {
		err := fmt.Errorf("SEAT NUMBER " + number + " ALREADY RESERVE!!")
		return err
	}
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
