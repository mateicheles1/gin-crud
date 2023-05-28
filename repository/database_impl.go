package repository

import "gorm.io/gorm"

type TodoListDBImpl struct {
	lists *gorm.DB
}

func NewTodoListDBImpl(db *gorm.DB) *TodoListDBImpl {
	return &TodoListDBImpl{
		lists: db,
	}
}
