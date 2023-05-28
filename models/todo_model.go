package models

type Todo struct {
	Id        string `gorm:"type:uuid;primary_key"`
	ListId    string `gorm:"type:uuid;not null"`
	Content   string `gorm:"not null"`
	Completed bool   `gorm:"not null"`
}
