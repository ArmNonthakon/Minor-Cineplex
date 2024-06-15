package adapters_http

import (
	"time"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/theater"
	"github.com/ArmNonthakon/Minor-Cineplex/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TheaterServiceIml struct {
	service ports.TheaterService
}

func NewTheaterHttpHandler(service ports.TheaterService) *TheaterServiceIml {
	return &TheaterServiceIml{service: service}
}

// Add theater
func (h *TheaterServiceIml) AddTheater(c *fiber.Ctx) error {
	input := domain.InputTheater{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	if err := h.service.ReqToAddTheater(input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Aready add theater number" + input.TheaterNumber)
	}
	return c.Status(fiber.StatusCreated).JSON("Success add new theater number " + input.TheaterNumber + " !!")
}

// Get theater
func (h *TheaterServiceIml) GetTheater(c *fiber.Ctx) error {
	data, err := h.service.ReqToGetTheater()
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	resultData := utils.TheaterTranToJson(data)
	return c.Status(fiber.StatusFound).JSON(resultData)
}

// Get theater sort by number
func (h *TheaterServiceIml) GetTheaterByNumber(c *fiber.Ctx) error {
	input := domain.InputNumber{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	data, err := h.service.ReqToGetTheaterByNumber(input.TheaterNumber)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if len(data) == 0 {
		return c.SendStatus(fiber.StatusNoContent)
	}
	resultData := utils.TheaterTranToJson(data)
	return c.Status(fiber.StatusFound).JSON(resultData)
}

// Get theater sort by id
func (h *TheaterServiceIml) GetTheaterById(c *fiber.Ctx) error {
	input := domain.InputTheaterId{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	theater, err := h.service.ReqToGetTheaterById(uuid.MustParse(input.TheaterId))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
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
	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"TheaterId":          theater.TheaterId,
		"TheaterNumber":      theater.TheaterNumber,
		"TimeStart":          theater.StartTime,
		"TimeEnd":            theater.StartTime.Add(time.Hour * time.Duration(theater.Movie.Duration)),
		"SeatMax":            theater.SeatCol * theater.SeatRow,
		"SeatExist":          (theater.SeatCol * theater.SeatRow) - len(theater.Seats),
		"Movies":             processedMovie,
		"SeatAlreadyReserve": processSeats,
	})
}
