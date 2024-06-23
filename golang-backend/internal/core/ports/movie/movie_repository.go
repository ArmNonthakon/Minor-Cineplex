package ports

import (
	"time"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

type MovieRepository interface {
	ResAllMovie() ([]domain.OnlyMovie, error)
	ResAllMovieWithTheater() ([]domain.Movie, error)
	ResMovieByTitle(title string) (domain.Movie, error)
	CreateMovie(title string, duration int, id uuid.UUID, genre string, date time.Time, url string) error
	DeleteMovieByTitle(title string) error
}
