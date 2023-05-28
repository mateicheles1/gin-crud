package models

type TodoList struct {
	Id        string  `gorm:"primary_key"`
	UserId    string  `gorm:"not null"`
	Owner     string  `gorm:"not null"`
	Completed bool    `gorm:"not null"`
	Todos     []*Todo `gorm:"foreignKey:ListId;constraint:onDelete:CASCADE"`
}
