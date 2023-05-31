package models

type JWT struct {
	UserId string `gorm:"primary_key"`
	Token  string `gorm:"not null"`
}
