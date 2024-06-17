package core

import (
	"fmt"

	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/user"
	"github.com/ArmNonthakon/Minor-Cineplex/utils"
)

type UserRepositoryIml struct {
	repo ports.UserRepository
}

func NewUserRepository(repo ports.UserRepository) ports.UserService {
	return &UserRepositoryIml{repo: repo}
}

func (r UserRepositoryIml) GetInputToLogin(userName string, password string) error {
	hashPass, err := r.repo.CheckLogin(userName)
	if err != nil {
		return err
	}
	checkPass := false
	if result := utils.CheckPasswordHash(password, hashPass); result == checkPass {
		err := fmt.Errorf("PASSWORD IS WRONG")
		return err
	}
	return nil
}
func (r UserRepositoryIml) GetInputToRegister(userName string, email string, password string) error {
	hashPass, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	if err := r.repo.CreateUser(userName, email, hashPass); err != nil {
		return err
	}
	return nil
}
