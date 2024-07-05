package main

import (
	"log"
	"os"

	adapters_gorm "github.com/ArmNonthakon/Minor-Cineplex/internal/adapters/gorm_adapter"
	adapters_http "github.com/ArmNonthakon/Minor-Cineplex/internal/adapters/http_adapter"
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core"
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	jwtware "github.com/gofiber/contrib/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(server string) (*gorm.DB, error) {
	var dsn string
	if server == "mysql" {
		dsn := "root:Ninjaarm-2003@tcp(127.0.0.1:3306)/minorcineplex?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		return DB, err
	} else {
		dsn = os.Getenv("postgres_database")
		DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		return DB, err
	}

}
func NewHttpHandler(db *gorm.DB) (adapters_http.MovieServiceIml, adapters_http.TheaterServiceIml, adapters_http.TicketServiceIml, adapters_http.UserServiceIml) {
	movieRepository := adapters_gorm.NewMovieGorm(db)
	movieService := core.NewMovieRepository(movieRepository)
	movieHttpHandler := adapters_http.NewMovieHttpHandler(movieService)

	theaterRepository := adapters_gorm.NewTheaterGorm(db)
	theaterService := core.NewTheaterRepository(theaterRepository)
	theaterHttpHandler := adapters_http.NewTheaterHttpHandler(theaterService)

	ticketRepository := adapters_gorm.NewTicketGorm(db)
	ticketService := core.NewTicketRepository(ticketRepository)
	ticketHttpHandler := adapters_http.NewTicketHttpHandler(ticketService)

	userRepository := adapters_gorm.NewUserGorm(db)
	userService := core.NewUserRepository(userRepository)
	userHttpHandler := adapters_http.NewUserService(userService)
	return *movieHttpHandler, *theaterHttpHandler, *ticketHttpHandler, *userHttpHandler
}
func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		/*AllowOrigins:     "http://localhost:5173/",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,*/
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	db, err := NewDB("postgres")
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&domain.User{}, &domain.Movie{}, &domain.Theater{}, &domain.Ticket{}, &domain.Seat{})
	if err != nil {
		log.Panic(err)
	}
	log.Println("Tables migrated successfully")
	movie, theater, ticket, user := NewHttpHandler(db)
	app.Get("/getMovie", movie.GetAllMovie)
	app.Get("/getMovieWithTheater", movie.GetAllMovieWithTheater)
	app.Post("/getMovieByTitle", movie.GetMovieByTitle)
	app.Post("/addMovie", movie.AddNewMovie)
	app.Delete("/deleteMovieByTitle", movie.DeleteMovieByTitle)

	app.Get("/getTheater", theater.GetTheater)
	app.Post("/getTheaterByNumber", theater.GetTheaterByNumber)
	app.Post("/addTheater", theater.AddTheater)
	app.Post("/getTheaterById", theater.GetTheaterById)
	app.Delete("/deleteTheater", theater.DeleteTheater)

	app.Post("/register", user.Register)
	app.Post("/login", user.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte("MinorCineplex")},
		TokenLookup: "cookie:token",
	}))
	if err != nil {
		log.Panic(err)
	}
	app.Post("/reserveSeat", ticket.ReserveTicket)
	app.Post("/getTicketById", ticket.GetTicketById)
	app.Listen(":3000")

}
