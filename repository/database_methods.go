package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mateicheles1/golang-crud/data"
	"github.com/mateicheles1/golang-crud/models"
	"gorm.io/gorm"
)

func (db *TodoListDBImpl) CreateList(listModel *models.TodoList) (*models.TodoList, error) {
	listModel.Id = uuid.New().String()

	for i := range listModel.Todos {
		listModel.Todos[i].Id = uuid.New().String()
		listModel.Todos[i].ListId = listModel.Id
	}

	err := db.lists.Create(listModel).Error

	if err != nil {
		return nil, err
	}

	return listModel, nil
}

func (db *TodoListDBImpl) CreateTodo(todoModel *models.Todo, listId string) (*models.Todo, error) {

	var list models.TodoList

	err := db.lists.First(&list, "id = ?", listId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	todoModel.Id = uuid.New().String()

	errCreate := db.lists.Create(todoModel).Error

	if errCreate != nil {
		return nil, err
	}

	return todoModel, nil
}

func (db *TodoListDBImpl) GetUser(userId string) (*models.User, error) {
	var user models.User

	err := db.lists.First(&user, "id = ?", userId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (db *TodoListDBImpl) GetList(listId string) (*models.TodoList, error) {
	var list models.TodoList

	err := db.lists.Preload("Todos").First(&list, "id = ?", listId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &list, nil
}

func (db *TodoListDBImpl) GetTodo(todoId string) (*models.Todo, error) {
	var todo models.Todo

	err := db.lists.First(&todo, "id = ?", todoId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &todo, nil
}

func (db *TodoListDBImpl) PatchList(completed bool, listId string) (*models.TodoList, error) {

	var list models.TodoList

	tx := db.lists.Begin()

	if err := tx.Table("todo_lists").Where("id = ?", listId).Update("completed", completed).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	if err := tx.Table("todo_lists").Where("id = ?", listId).Preload("Todos").First(&list).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &list, nil
}

func (db *TodoListDBImpl) PatchTodo(completed bool, todoId string) (*models.Todo, error) {
	var todo models.Todo

	tx := db.lists.Begin()

	if err := tx.Table("todos").Where("id = ?", todoId).Update("completed", completed).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	if err := tx.Table("todos").Where("id = ?", todoId).First(&todo).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &todo, nil
}

func (db *TodoListDBImpl) DeleteList(listId string) error {
	err := db.lists.Table("todo_lists").Where("id = ?", listId).Delete(&models.TodoList{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return err
	}

	return nil
}

func (db *TodoListDBImpl) DeleteTodo(todoId string) error {
	err := db.lists.Table("todos").Where("id = ?", todoId).Delete(&models.Todo{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return err
	}

	return nil
}

func (db *TodoListDBImpl) GetLists(userId string, completedBool *bool) ([]*models.TodoList, error) {
	var user models.User

	err := db.lists.Preload("Todolists.Todos").First(&user, "id = ?", userId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	if completedBool != nil {

		var filteredList []*models.TodoList

		for _, list := range user.Todolists {

			if list.Completed == *completedBool {
				filteredList = append(filteredList, list)
			}

		}

		return filteredList, nil
	}

	return user.Todolists, nil
}

func (db *TodoListDBImpl) CreateUser(userModel *models.User) (*models.User, error) {
	userModel.Id = uuid.New().String()

	err := db.lists.Create(userModel).Error

	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (db *TodoListDBImpl) Login(reqBody *data.UserLoginDTO) (*models.User, error) {
	var user models.User

	err := db.lists.First(&user, "username = ? AND password = ?", reqBody.Username, reqBody.Password).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (db *TodoListDBImpl) GetToken(userId string) (*models.JWT, error) {
	var token models.JWT

	if err := db.lists.First(&token, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	return &token, nil

}

func (db *TodoListDBImpl) StoreJWT(jwt *models.JWT) error {

	if err := db.lists.Create(jwt).Error; err != nil {
		return err
	}

	return nil
}
