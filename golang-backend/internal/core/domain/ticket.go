package domain

type Ticket struct {
	TicketId string `gorm:"primaryKey"`
	Status   string `gorm:"not null"`
	UserID   string
	Users    Users `gorm:"references:UserID"`
}

type TicketSeat struct {
	TicketId   string
	Ticket     Ticket `gorm:"references:TicketId"`
	SeatNumber string
	Seat       Seat `gorm:"references:SeatNumber"`
}
