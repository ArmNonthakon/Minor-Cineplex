package ports

type UserRepository interface {
	CheckLogin(UserName string) (string, error)
	CreateUser(userName string, email string, password string) error
}
