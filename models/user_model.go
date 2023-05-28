package models

type User struct {
	Id        string      `gorm:"not null"`
	Username  string      `gorm:"primary_key;not null"`
	Password  string      `gorm:"not null"`
	Role      string      `gorm:"not null"`
	Todolists []*TodoList `gorm:"foreignKey:Owner;constraint:onDelete:CASCADE,onUpdate:CASCADE"`
}
