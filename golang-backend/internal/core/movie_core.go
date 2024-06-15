package core

import (
	"strings"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/movie"
	"github.com/google/uuid"
)

// Other Func

// Main Func
type MovieRepositoryIml struct {
	repo ports.MovieRepository
}

func NewMovieRepository(repo ports.MovieRepository) ports.MovieService {
	return &MovieRepositoryIml{repo: repo}
}

// Get movie
func (r MovieRepositoryIml) ReqForAllMovie() ([]domain.OnlyMovie, error) {
	data, err := r.repo.ResAllMovie()
	if err != nil {
		return data, err
	}
	return data, nil
}

// Get moive with theater
func (r MovieRepositoryIml) ReqForAllMovieWithTheater() ([]domain.Movie, error) {
	data, err := r.repo.ResAllMovieWithTheater()
	if err != nil {
		return data, err
	}
	return data, nil
}

// Get movie sort by title
func (r MovieRepositoryIml) ReqForMovieBYTitle(title string) (domain.Movie, error) {
	data, err := r.repo.ResMovieByTitle(strings.ToLower(title))
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add new movie
func (r MovieRepositoryIml) ReqToAddMovie(input domain.InputMovie) error {
	err := r.repo.CreateMovie(strings.ToLower(input.Title), input.Duration, uuid.New())
	if err != nil {
		return err
	}
	return nil
}
