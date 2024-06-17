package adapters_gorm

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/movie"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewMovieGorm(db *gorm.DB) ports.MovieRepository {
	return &GormDB{db: db}
}

// Get movie
func (g *GormDB) ResAllMovie() ([]domain.OnlyMovie, error) {
	data := []domain.OnlyMovie{}
	if result := g.db.Select("movie_id,title").Model(&domain.Movie{}).Find(&data); result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

// Get movie with theater
func (g *GormDB) ResAllMovieWithTheater() ([]domain.Movie, error) {
	data := []domain.Movie{}
	if result := g.db.Model(&domain.Movie{}).Preload("Theaters").Find(&data); result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

// Get movie sort by title
func (g *GormDB) ResMovieByTitle(title string) (domain.Movie, error) {
	data := domain.Movie{}
	if result := g.db.Model(&domain.Movie{}).Preload("Theaters").Find(&data, "title = ?", title); result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

// Add movie
func (g *GormDB) CreateMovie(title string, duration float32, id uuid.UUID) error {
	data := domain.Movie{
		MovieId:  id,
		Title:    title,
		Duration: duration,
	}
	if result := g.db.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete movie by title
func (g *GormDB) DeleteMovieByTitle(title string) error {
	var movie domain.Movie
	if err := g.db.Where("title = ?", title).Preload("Theaters").First(&movie).Error; err != nil {
		return err
	}
	if err := g.db.Model(&movie).Association("Theaters").Clear(); err != nil {
		return err
	}
	if err := g.db.Model(&movie).Association("Tickets").Clear(); err != nil {
		return err
	}
	if err := g.db.Delete(&movie).Error; err != nil {
		return err
	}

	return nil
}
