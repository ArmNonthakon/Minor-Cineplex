package adapters_http

import (
	"log"
	"time"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/movie"
	"github.com/ArmNonthakon/Minor-Cineplex/utils"
	"github.com/gofiber/fiber/v2"
)

type MovieServiceIml struct {
	service ports.MovieService
}

func NewMovieHttpHandler(service ports.MovieService) *MovieServiceIml {
	return &MovieServiceIml{service: service}
}

// Get movie
func (h *MovieServiceIml) GetAllMovie(c *fiber.Ctx) error {
	data, err := h.service.ReqForAllMovie()
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusAccepted).JSON(data)
}

// Get movie with theater
func (h *MovieServiceIml) GetAllMovieWithTheater(c *fiber.Ctx) error {
	data, err := h.service.ReqForAllMovieWithTheater()
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	resultData := utils.MovieTranToJson(data)
	return c.Status(fiber.StatusAccepted).JSON(resultData)
}

// Get movie sort by title
func (h *MovieServiceIml) GetMovieByTitle(c *fiber.Ctx) error {
	input := domain.InputTitle{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	data, err := h.service.ReqForMovieBYTitle(input.Title)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	theaterData := make([]map[string]interface{}, len(data.Theaters))
	for j, theater := range data.Theaters {
		theaterData[j] = map[string]interface{}{
			"TheaterId":     theater.TheaterId,
			"TheaterNumber": theater.TheaterNumber,
			"MaxSeat":       theater.SeatCol * theater.SeatRow,
			"TimeStart":     theater.StartTime,
			"SeatCol":       theater.SeatCol,
			"SeatRow":       theater.SeatRow,
			"TimeEnd":       theater.StartTime.Add(time.Minute * time.Duration(data.Duration)),
		}
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"MovieId":      data.MovieId,
		"MovieTitle":   data.Title,
		"MovieGenre":   data.Genre,
		"ReleaseDate":  data.ReleaseDate,
		"Poster":       data.PosterUrl,
		"TimeDuration": data.Duration,
		"Theaters":     theaterData,
	})
}

// Add new movie
func (h *MovieServiceIml) AddNewMovie(c *fiber.Ctx) error {
	input := domain.InputMovie{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	err := h.service.ReqToAddMovie(input)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("Aready add movie " + input.Title)
	}
	return c.Status(fiber.StatusCreated).JSON("Success add new movie " + input.Title + " !!")
}

// Delete movie by title
func (h *MovieServiceIml) DeleteMovieByTitle(c *fiber.Ctx) error {
	input := domain.InputTitle{}
	if err := c.BodyParser(&input); err != nil {
		log.Panic(err)
		return c.Status(fiber.StatusNotAcceptable).JSON("Error Input")
	}
	err := h.service.ReqToDeleteMovieByTitle(input.Title)
	if err != nil {
		log.Panic(err)
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON("DELETE SUCCESS!!")
}
