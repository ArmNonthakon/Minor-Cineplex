package adapters_http

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/ticker"
	"github.com/ArmNonthakon/Minor-Cineplex/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Main func
type TicketServiceIml struct {
	service ports.TicketService
}

func NewTicketHttpHandler(service ports.TicketService) *TicketServiceIml {
	return &TicketServiceIml{service: service}
}

// Reserve seat and recieve ticket
func (h *TicketServiceIml) ReserveTicket(c *fiber.Ctx) error {
	input := domain.InputTicket{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	ticket, err := h.service.ReqReserveSeat(input)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Seats have already reserved or incorrect input!!")
	}
	resultSeat := utils.SliceSeat(ticket)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"TicketId":      ticket.TicketId,
		"UserName":      ticket.User.UserName,
		"MovieTitle":    ticket.Movie.Title,
		"Seats":         resultSeat,
		"TheaterNumber": ticket.Seats[0].Theater.TheaterNumber,
	})
}

// Get ticket by id
func (h *TicketServiceIml) GetTicketById(c *fiber.Ctx) error {
	input := domain.InputTicketId{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	ticket, err := h.service.ReqGetTicketById(input.TicketId)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	checkUUID := uuid.UUID{}
	if ticket.TicketId == checkUUID {
		return c.SendStatus(fiber.StatusNotFound)
	}
	resultSeat := utils.SliceSeat(ticket)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"TicketId":      ticket.TicketId,
		"UserName":      ticket.User.UserName,
		"MovieTitle":    ticket.Movie.Title,
		"Seats":         resultSeat,
		"TheaterNumber": ticket.Seats[0].Theater.TheaterNumber,
		"ShowTime":      ticket.Seats[0].Theater.StartTime,
	})
}
func (h *TicketServiceIml) GetTicketByUserName(c *fiber.Ctx) error {
	input := domain.User{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	ticket, err := h.service.ReqGetTicketByUserName(input.UserName)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	processedData := make([]map[string]interface{}, len(ticket))
	for i, data := range ticket {
		resultSeat := utils.SliceSeat(data)
		processedData[i] = map[string]interface{}{
			"TicketId":      data.TicketId,
			"UserName":      data.User.UserName,
			"MovieTitle":    data.Movie.Title,
			"Seats":         resultSeat,
			"TheaterNumber": data.Seats[0].Theater.TheaterNumber,
			"ShowTime":      data.Seats[0].Theater.StartTime,
		}
	}
	return c.Status(fiber.StatusAccepted).JSON(processedData)
}
