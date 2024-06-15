package ports

import "github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"

type MovieService interface {
	ReqForAllMovie() ([]domain.OnlyMovie, error)
	ReqForAllMovieWithTheater() ([]domain.Movie, error)
	ReqForMovieBYTitle(title string) (domain.Movie, error)
	ReqToAddMovie(input domain.InputMovie) error
}
