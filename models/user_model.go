package models

type User struct {
	Id        string      `gorm:"primary_key"`
	Username  string      `gorm:"not null"`
	Password  string      `gorm:"not null"`
	Role      string      `gorm:"not null"`
	Todolists []*TodoList `gorm:"foreignKey:UserId;constraint:onDelete:CASCADE,onUpdate:CASCADE"`
}
