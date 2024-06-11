package domain

type Seat struct {
	SeatNumber string `gorm:"primaryKey"`
	IsReserve  bool
	Status     string
	TheaterId  string
	Theater    Theater `gorm:"references:TheaterID"`
}
