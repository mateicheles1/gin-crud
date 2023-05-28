package repository

import (
	"github.com/mateicheles1/golang-crud/data"
	"github.com/mateicheles1/golang-crud/models"
)

type TodoListDB interface {
	CreateList(listModel *models.TodoList) (*models.TodoList, error)
	CreateTodo(todoModel *models.Todo, listId string) (*models.Todo, error)
	GetUser(userId string) (*models.User, error)
	GetList(listId string) (*models.TodoList, error)
	GetTodo(todoId string) (*models.Todo, error)
	PatchList(owner string, listId string, user *models.User) (*models.TodoList, error)
	PatchTodo(completed bool, todoId string) (*models.Todo, error)
	DeleteList(listId string) error
	DeleteTodo(todoId string) error
	GetLists(userId string) ([]*models.TodoList, error)
	CreateUser(userModel *models.User) (*models.User, error)
	Login(reqBody *data.UserLoginDTO) (*models.User, error)
}
