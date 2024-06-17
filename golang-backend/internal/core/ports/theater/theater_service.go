package ports

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

type TheaterService interface {
	ReqToAddTheater(input domain.InputTheater) error
	ReqToGetTheater() ([]domain.Theater, error)
	ReqToGetTheaterByNumber(number string) ([]domain.Theater, error)
	ReqToGetTheaterById(id uuid.UUID) (domain.Theater, error)
	ReqToDeleteTheater() error
}
