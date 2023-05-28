package service

import "github.com/mateicheles1/golang-crud/repository"

type TodoListServiceImpl struct {
	db repository.TodoListDB
}

func NewTodolistServiceImpl(db repository.TodoListDB) *TodoListServiceImpl {
	return &TodoListServiceImpl{
		db: db,
	}
}
