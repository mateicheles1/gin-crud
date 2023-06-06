package repository

import "gorm.io/gorm"

type TodoListDBImpl struct {
	lists *gorm.DB
}

func NewTodoListDB(db *gorm.DB) TodoListDB {
	return &TodoListDBImpl{
		lists: db,
	}
}
