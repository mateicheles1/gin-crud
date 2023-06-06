package service

import "github.com/mateicheles1/golang-crud/repository"

type TodoListServiceImpl struct {
	db repository.TodoListDB
}

func NewTodolistService(db repository.TodoListDB) TodoListService {
	return &TodoListServiceImpl{
		db: db,
	}
}
