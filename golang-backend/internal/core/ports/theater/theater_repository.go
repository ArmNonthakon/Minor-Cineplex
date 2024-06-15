package ports

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	"github.com/google/uuid"
)

type TheaterRepository interface {
	CreateTheater(data domain.Theater) error
	ResTheater() ([]domain.Theater, error)
	ResTheaterByNumber(number string) ([]domain.Theater, error)
	ResTheaterById(id uuid.UUID) (domain.Theater, error)
}
