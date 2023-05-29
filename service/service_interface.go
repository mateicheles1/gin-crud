package service

import "github.com/mateicheles1/golang-crud/data"

type TodoListService interface {
	CreateList(reqBody *data.TodoListCreateRequestDTO, userId string) (*data.TodoListResourceResponseDTO, error)
	CreateTodo(reqBody *data.TodoCreateRequestDTO, userId string, listId string) (*data.TodoResourceResponseDTO, error)
	GetList(listId string, username string) (*data.TodoListGetResponseDTO, error)
	GetTodo(todoId string, userId string) (*data.TodoGetResponseDTO, error)
	PatchList(reqBody *data.TodoListPatchDTO, listId string, username string) (*data.TodoListResourceResponseDTO, error)
	PatchTodo(reqBody *data.TodoPatchDTO, todoId string, username string) (*data.TodoResourceResponseDTO, error)
	DeleteList(listId string, username string) error
	DeleteTodo(todoId string, username string) error
	GetLists(userId string, completedBool *bool) (*[]*data.TodoListResourceResponseDTO, error)
	GetTodos(listId string, username string, completedBool *bool) (*[]*data.TodoGetResponseInListDTO, error)
	CreateUser(reqBody *data.UserCreateDTO) (*data.UserResponseDTO, error)
	Login(reqBody *data.UserLoginDTO) (string, error)
}
