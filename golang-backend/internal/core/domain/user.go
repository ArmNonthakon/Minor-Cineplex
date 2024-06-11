package domain

type Users struct {
	UserID   string `gorm:"primarykey"`
	UserName string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Password string
}
