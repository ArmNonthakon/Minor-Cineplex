package main

import (
	"log"

	adapters_gorm "github.com/ArmNonthakon/Minor-Cineplex/internal/adapters/gorm_adapter"
	adapters_http "github.com/ArmNonthakon/Minor-Cineplex/internal/adapters/http_adapter"
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core"
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "root:Ninjaarm-2003@tcp(127.0.0.1:3306)/MinorCineplex?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return DB, err
}
func NewHttpHandler(db *gorm.DB) (adapters_http.MovieServiceIml, adapters_http.TheaterServiceIml, adapters_http.TicketServiceIml) {
	movieRepository := adapters_gorm.NewMovieGorm(db)
	movieService := core.NewMovieRepository(movieRepository)
	movieHttpHandler := adapters_http.NewMovieHttpHandler(movieService)

	theaterRepository := adapters_gorm.NewTheaterGorm(db)
	theaterService := core.NewTheaterRepository(theaterRepository)
	theaterHttpHandler := adapters_http.NewTheaterHttpHandler(theaterService)

	ticketRepository := adapters_gorm.NewTicketGorm(db)
	ticketService := core.NewTicketRepository(ticketRepository)
	ticketHttpHandler := adapters_http.NewTicketHttpHandler(ticketService)
	return *movieHttpHandler, *theaterHttpHandler, *ticketHttpHandler
}
func main() {
	app := fiber.New()
	db, err := NewDB()
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&domain.User{}, &domain.Movie{}, &domain.Theater{}, &domain.Ticket{}, &domain.Seat{})
	if err != nil {
		log.Panic(err)
	}
	log.Println("Tables migrated successfully")
	movie, theater, ticket := NewHttpHandler(db)
	app.Get("/getMovie", movie.GetAllMovie)
	app.Get("/getMovieWithTheater", movie.GetAllMovieWithTheater)
	app.Post("/getMovieByTitle", movie.GetMovieByTitle)
	app.Post("/addMovie", movie.AddNewMovie)

	app.Get("/getTheater", theater.GetTheater)
	app.Post("/getTheaterByNumber", theater.GetTheaterByNumber)
	app.Post("/addTheater", theater.AddTheater)
	app.Post("/getTheaterById", theater.GetTheaterById)

	app.Post("/reserveSeat", ticket.ReserveTicket)
	app.Post("/getTicketById", ticket.GetTicketById)
	log.Fatal(app.Listen(":3000"))
	/*id := uuid.MustParse("bcb93213-8bc7-416d-a7d5-a59c74e41bfb")
	var movie domain.Movie
	db.Preload("Theaters").First(&movie, id) // Assume the movie with ID 1 exists

	// Clear associations with theaters
	db.Model(&movie).Association("Theaters").Clear()*/

}
