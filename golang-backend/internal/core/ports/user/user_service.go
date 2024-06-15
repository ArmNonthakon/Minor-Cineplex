package ports

type UserService interface {
	GetInputToLogin(userName string, password string) error
	GetInputToRegister(userName string, email string, password string) error
}
