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
		return c.Status(fiber.StatusNotAcceptable).JSON("Aready reserve")
	}
	resultSeat := utils.SliceSeat(ticket)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"TicketId":      ticket.TicketId,
		"UserId":        ticket.UserId,
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
		"UserId":        ticket.UserId,
		"MovieTitle":    ticket.Movie.Title,
		"Seats":         resultSeat,
		"TheaterNumber": ticket.Seats[0].Theater.TheaterNumber,
	})
}
