package adapters_gorm

import (
	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/user"
	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) ports.UserRepository {
	return &GormDB{db: db}
}
func (g *GormDB) CheckLogin(userName string) (string, error) {
	data := domain.User{}
	if result := g.db.Select("password").Find(&data, "userName= ?", userName); result.Error != nil {
		return "", result.Error
	}
	return data.Password, nil
}
func (g *GormDB) CreateUser(userName string, email string, password string) error {
	data := domain.User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
	if result := g.db.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}
