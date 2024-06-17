package adapters_gorm

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/theater"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewTheaterGorm(db *gorm.DB) ports.TheaterRepository {
	return &GormDB{db: db}
}

// Add theater
func (g *GormDB) CreateTheater(data domain.Theater) error {
	if result := g.db.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}

// Get theater
func (g *GormDB) ResTheater() ([]domain.Theater, error) {
	data := []domain.Theater{}
	if result := g.db.Preload("Movie").Preload("Seats").Find(&data); result.Error != nil {
		return data, result.Error
	}
	return data, nil
}
func (g *GormDB) ResTheaterByNumber(number string) ([]domain.Theater, error) {
	data := []domain.Theater{}
	if result := g.db.Model(&domain.Theater{}).Preload("Movie").Preload("Seats").Find(&data, "theater_number = ?", number); result.Error != nil {
		return data, result.Error
	}
	return data, nil
}
func (g *GormDB) ResTheaterById(id uuid.UUID) (domain.Theater, error) {
	data := domain.Theater{}
	if result := g.db.Model(&domain.Theater{}).Preload("Movie").Preload("Seats").Find(&data, id); result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

// Delete Theater
func (g *GormDB) DeleteTheater() error {
	var theater domain.Theater
	if err := g.db.First(&theater).Error; err != nil {
		return err
	}
	if err := g.db.Model(&theater).Association("Seats").Clear(); err != nil {
		return err
	}
	if err := g.db.Delete(&theater).Error; err != nil {
		return err
	}

	return nil
}
