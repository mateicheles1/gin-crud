package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mateicheles1/golang-crud/data"
	"github.com/mateicheles1/golang-crud/models"
)

func (s *TodoListServiceImpl) CreateList(reqBody *data.TodoListCreateRequestDTO, userId string) (*data.TodoListResourceResponseDTO, error) {

	user, err := s.db.GetUser(userId)

	if err != nil {
		return nil, err
	}

	listModel := models.TodoList{
		Owner:     user.Username,
		UserId:    user.Id,
		Completed: false,
		Todos:     make([]*models.Todo, len(reqBody.Todos)),
	}

	for i, v := range reqBody.Todos {
		listModel.Todos[i] = &models.Todo{
			Content:   v,
			Completed: false,
		}
	}

	list, err := s.db.CreateList(&listModel)

	if err != nil {
		return nil, err
	}

	todoListResource := data.TodoListResourceResponseDTO{
		Id:        list.Id,
		UserId:    user.Id,
		Owner:     list.Owner,
		Completed: list.Completed,
		Todos:     make([]*data.TodoResourceResponseDTO, len(list.Todos)),
	}

	for i := range list.Todos {
		todoListResource.Todos[i] = &data.TodoResourceResponseDTO{
			Id:        list.Todos[i].Id,
			ListId:    list.Todos[i].ListId,
			Content:   list.Todos[i].Content,
			Completed: list.Todos[i].Completed,
		}
	}

	return &todoListResource, nil

}

func (s *TodoListServiceImpl) CreateTodo(reqBody *data.TodoCreateRequestDTO, userId string, listId string) (*data.TodoResourceResponseDTO, error) {

	todoModel := models.Todo{
		ListId:    listId,
		Content:   reqBody.Content,
		Completed: false,
	}

	user, err := s.db.GetUser(userId)

	if err != nil {
		return nil, err
	}

	list, err := s.db.GetList(listId)

	if err != nil {
		return nil, err
	}

	if list.Owner != user.Username {
		return nil, errors.New("action not allowed")
	}

	todo, err := s.db.CreateTodo(&todoModel, listId)

	if err != nil {
		return nil, err
	}

	todoResponse := data.TodoResourceResponseDTO{
		Id:        todo.Id,
		ListId:    todo.ListId,
		Content:   todo.Content,
		Completed: todo.Completed,
	}

	return &todoResponse, nil

}

func (s *TodoListServiceImpl) GetList(listId string, username string) (*data.TodoListGetResponseDTO, error) {
	list, err := s.db.GetList(listId)

	if err != nil {
		return nil, err
	}

	if list.Owner != username {
		return nil, errors.New("action not allowed")
	}

	listResponse := data.TodoListGetResponseDTO{
		Owner:     list.Owner,
		Completed: list.Completed,
		Todos:     make([]*data.TodoGetResponseInListDTO, len(list.Todos)),
	}

	for i := range list.Todos {
		listResponse.Todos[i] = &data.TodoGetResponseInListDTO{
			Id:        list.Todos[i].Id,
			Content:   list.Todos[i].Content,
			Completed: list.Todos[i].Completed,
		}
	}

	return &listResponse, nil
}

func (s *TodoListServiceImpl) GetTodo(todoId string, username string) (*data.TodoGetResponseDTO, error) {

	todo, err := s.db.GetTodo(todoId)

	if err != nil {
		return nil, err
	}

	list, err := s.db.GetList(todo.ListId)

	if err != nil {
		return nil, err
	}

	if list.Owner != username {
		return nil, errors.New("action not allowed")
	}

	todoResponse := data.TodoGetResponseDTO{
		ListId:    todo.ListId,
		Content:   todo.Content,
		Completed: todo.Completed,
	}

	return &todoResponse, nil

}

func (s *TodoListServiceImpl) PatchList(reqBody *data.TodoListPatchDTO, listId string, userId string) (*data.TodoListResourceResponseDTO, error) {
	list, err := s.db.GetList(listId)

	if err != nil {
		return nil, err
	}

	user, err := s.db.GetUser(userId)

	if err != nil {
		return nil, err
	}

	if list.Owner != user.Username {
		return nil, errors.New("action not allowed")
	}

	patchedList, err := s.db.PatchList(reqBody.Completed, listId)

	if err != nil {
		return nil, err
	}

	responseList := data.TodoListResourceResponseDTO{
		Id:        patchedList.Id,
		UserId:    patchedList.UserId,
		Owner:     patchedList.Owner,
		Completed: patchedList.Completed,
		Todos:     make([]*data.TodoResourceResponseDTO, len(patchedList.Todos)),
	}

	for i := range patchedList.Todos {
		responseTodos := data.TodoResourceResponseDTO{
			Id:        patchedList.Todos[i].Id,
			ListId:    patchedList.Todos[i].ListId,
			Content:   patchedList.Todos[i].Content,
			Completed: patchedList.Todos[i].Completed,
		}

		responseList.Todos[i] = &responseTodos
	}

	return &responseList, nil

}

func (s *TodoListServiceImpl) PatchTodo(reqBody *data.TodoPatchDTO, todoId string, username string) (*data.TodoResourceResponseDTO, error) {
	todo, err := s.db.GetTodo(todoId)

	if err != nil {
		return nil, err
	}

	list, err := s.db.GetList(todo.ListId)

	if err != nil {
		return nil, err
	}

	if list.Owner != username {
		return nil, errors.New("action not allowed")
	}

	patchedTodo, err := s.db.PatchTodo(reqBody.Completed, todoId)

	if err != nil {
		return nil, err
	}

	responseTodo := data.TodoResourceResponseDTO{
		Id:        patchedTodo.Id,
		ListId:    patchedTodo.ListId,
		Content:   patchedTodo.Content,
		Completed: patchedTodo.Completed,
	}

	return &responseTodo, nil
}

func (s *TodoListServiceImpl) DeleteList(listId string, username string) error {
	list, err := s.db.GetList(listId)

	if err != nil {
		return err
	}

	if list.Owner != username {
		return errors.New("action not allowed")
	}

	errTwo := s.db.DeleteList(listId)

	if errTwo != nil {
		return errTwo
	}

	return nil
}

func (s *TodoListServiceImpl) DeleteTodo(todoId string, username string) error {
	todo, err := s.db.GetTodo(todoId)

	if err != nil {
		return err
	}

	list, err := s.db.GetList(todo.ListId)

	if err != nil {
		return err
	}

	if list.Owner != username {
		return errors.New("action not allowed")
	}

	errTodo := s.db.DeleteTodo(todoId)

	if errTodo != nil {
		return errTodo
	}

	return nil
}

func (s *TodoListServiceImpl) GetLists(userId string, completedBool *bool) (*[]*data.TodoListResourceResponseDTO, error) {

	lists, err := s.db.GetLists(userId, completedBool)

	if err != nil {
		return nil, err
	}

	listsResponse := make([]*data.TodoListResourceResponseDTO, len(lists))

	for i := range lists {

		listsResponse[i] = &data.TodoListResourceResponseDTO{
			Id:        lists[i].Id,
			Owner:     lists[i].Owner,
			Completed: lists[i].Completed,
			Todos:     make([]*data.TodoResourceResponseDTO, len(lists[i].Todos)),
		}

		for j := range lists[i].Todos {
			listsResponse[i].Todos[j] = &data.TodoResourceResponseDTO{
				Id:        lists[i].Todos[j].Id,
				Content:   lists[i].Todos[j].Content,
				Completed: lists[i].Todos[j].Completed,
			}
		}

	}

	return &listsResponse, nil
}

func (s *TodoListServiceImpl) GetTodos(listId string, username string, completedBool *bool) (*[]*data.TodoGetResponseInListDTO, error) {
	list, err := s.db.GetList(listId)

	if err != nil {
		return nil, err
	}

	if list.Owner != username {
		return nil, errors.New("action not allowed")
	}

	var responseTodo []*data.TodoGetResponseInListDTO

	if completedBool != nil {

		for i := range list.Todos {
			if list.Todos[i].Completed == *completedBool {
				todo := &data.TodoGetResponseInListDTO{
					Id:        list.Todos[i].Id,
					Content:   list.Todos[i].Content,
					Completed: list.Todos[i].Completed,
				}

				responseTodo = append(responseTodo, todo)
			}
		}

		return &responseTodo, nil
	}

	for i := range list.Todos {
		todo := &data.TodoGetResponseInListDTO{
			Id:        list.Todos[i].Id,
			Content:   list.Todos[i].Content,
			Completed: list.Todos[i].Completed,
		}

		responseTodo = append(responseTodo, todo)
	}

	return &responseTodo, nil
}

func (s *TodoListServiceImpl) CreateUser(reqBody *data.UserCreateDTO) (*data.UserResponseDTO, error) {
	userModel := models.User{
		Username: reqBody.Username,
		Password: reqBody.Password,
		Role:     reqBody.Role,
	}

	user, err := s.db.CreateUser(&userModel)

	if err != nil {
		return nil, err
	}

	userResponse := data.UserResponseDTO{
		Id:       user.Id,
		Username: user.Username,
	}

	return &userResponse, nil

}

func (s *TodoListServiceImpl) Login(reqBody *data.UserLoginDTO) (string, error) {
	user, err := s.db.Login(reqBody)

	if err != nil {
		return "", err
	}

	token, _ := s.db.GetToken(user.Id)

	if token != nil {
		return token.Token, nil
	}

	claims := jwt.MapClaims{
		"userId":   user.Id,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := newToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	DBtoken := models.JWT{
		UserId: user.Id,
		Token:  signedToken,
	}

	if tokenErr := s.db.StoreJWT(&DBtoken); tokenErr != nil {
		return "", err
	}

	return DBtoken.Token, nil

}
