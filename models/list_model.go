package models

type TodoList struct {
	Id    string  `gorm:"primary_key"`
	Owner string  `gorm:"not null"`
	Todos []*Todo `gorm:"foreignKey:ListId;constraint:onDelete:CASCADE"`
}
