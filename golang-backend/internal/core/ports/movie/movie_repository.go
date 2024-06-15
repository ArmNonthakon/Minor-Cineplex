package ports

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

type MovieRepository interface {
	ResAllMovie() ([]domain.OnlyMovie, error)
	ResAllMovieWithTheater() ([]domain.Movie, error)
	ResMovieByTitle(title string) (domain.Movie, error)
	CreateMovie(title string, duration float32, id uuid.UUID) error
}
