package core

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/theater"
	"github.com/google/uuid"
)

type TheaterRepositoryIml struct {
	repo ports.TheaterRepository
}

func NewTheaterRepository(repo ports.TheaterRepository) ports.TheaterService {
	return &TheaterRepositoryIml{repo: repo}
}

// Add new theater
func (r *TheaterRepositoryIml) ReqToAddTheater(input domain.InputTheater) error {
	data := domain.Theater{
		TheaterId:     uuid.New(),
		TheaterNumber: input.TheaterNumber,
		StartTime:     input.StartTime,
		MovieId:       input.MovieId,
		SeatRow:       input.SeatRow,
		SeatCol:       input.SeatCol,
	}
	if err := r.repo.CreateTheater(data); err != nil {
		return err
	}
	return nil
}

// Get theater
func (r *TheaterRepositoryIml) ReqToGetTheater() ([]domain.Theater, error) {
	data, err := r.repo.ResTheater()
	if err != nil {
		return data, err
	}
	return data, nil
}

// Get theater sort by theater number
func (r *TheaterRepositoryIml) ReqToGetTheaterByNumber(number string) ([]domain.Theater, error) {
	data, err := r.repo.ResTheaterByNumber(number)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (r *TheaterRepositoryIml) ReqToGetTheaterById(id uuid.UUID) (domain.Theater, error) {
	data, err := r.repo.ResTheaterById(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *TheaterRepositoryIml) ReqToDeleteTheater() error {
	err := r.repo.DeleteTheater()
	if err != nil {
		return err
	}
	return nil
}
