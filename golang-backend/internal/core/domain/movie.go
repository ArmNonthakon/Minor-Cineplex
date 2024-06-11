package domain

type Movie struct {
	MovieId string `gorm:"primaryKey"`
	Title   string
}
