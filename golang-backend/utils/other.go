package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
)

func SliceSeat(ticket domain.Ticket) []string {
	resultSeat := []string{}
	for _, seat := range ticket.Seats {
		resultSeat = append(resultSeat, seat.SeatNumber)

	}
	return resultSeat
}
func TheaterTranToJson(data []domain.Theater) []map[string]interface{} {
	processedData := make([]map[string]interface{}, len(data))
	for i, theater := range data {
		processedMovie := map[string]interface{}{
			"MovieId":       theater.MovieId,
			"MovieTitle":    theater.Movie.Title,
			"MovieDuration": theater.Movie.Duration,
		}
		processSeats := make([]map[string]interface{}, len(theater.Seats))
		for j, seat := range theater.Seats {
			processSeats[j] = map[string]interface{}{
				"SeatId":     seat.SeatID,
				"SeatNumber": seat.SeatNumber,
				"TicketId":   seat.TicketId,
			}
		}
		processedData[i] = map[string]interface{}{
			"TheaterId":          theater.TheaterId,
			"TheaterNumber":      theater.TheaterNumber,
			"TimeStart":          theater.StartTime,
			"TimeEnd":            theater.StartTime.Add(time.Hour * time.Duration(theater.Movie.Duration)),
			"SeatMax":            theater.SeatCol * theater.SeatRow,
			"SeatExist":          (theater.SeatCol * theater.SeatRow) - len(theater.Seats),
			"SeatCol":            theater.SeatCol,
			"SeatRow":            theater.SeatRow,
			"Movies":             processedMovie,
			"SeatAlreadyReserve": processSeats,
		}
	}
	return processedData
}
func MovieTranToJson(data []domain.Movie) []map[string]interface{} {
	processedData := make([]map[string]interface{}, len(data))
	for i, movie := range data {
		theaterData := make([]map[string]interface{}, len(movie.Theaters))
		for j, theater := range movie.Theaters {
			theaterData[j] = map[string]interface{}{
				"TheaterId":     theater.TheaterId,
				"TheaterNumber": theater.TheaterNumber,
				"MaxSeat":       theater.SeatCol * theater.SeatRow,
				"TimeStart":     theater.StartTime,
				"TimeEnd":       theater.StartTime.Add(time.Hour * time.Duration(movie.Duration)),
			}
		}
		processedData[i] = map[string]interface{}{
			"MovieId":      movie.MovieId,
			"MovieTitle":   movie.Title,
			"TimeDuration": movie.Duration,
			"Theaters":     theaterData,
		}
	}
	return processedData
}
func CalculateSeats(seatRow int, seatCol int, alreadyReserveSeat []string) []string {
	allSeat := []string{}
	for i := 0; i < seatRow; i++ {
		s := strconv.Itoa(65 + i)
		for j := 0; j < seatCol; j++ {
			allSeat = append(allSeat, fmt.Sprintf(s+"%d", j))
		}
	}
	return allSeat
}
