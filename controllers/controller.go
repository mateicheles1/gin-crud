package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateicheles1/golang-crud/data"
	"gorm.io/gorm"
)

func (c *Controller) CreateList(ctx *gin.Context) {
	var reqBody data.TodoListCreateRequestDTO

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	userId := ctx.MustGet("userId").(string)

	list, err := c.service.CreateList(&reqBody, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Location", fmt.Sprintf("api/v3/lists/%s", list.Id))

	ctx.JSON(http.StatusCreated, list)
}

func (c *Controller) CreateTodo(ctx *gin.Context) {
	var reqBody data.TodoCreateRequestDTO

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	userId := ctx.MustGet("userId").(string)

	todo, err := c.service.CreateTodo(&reqBody, userId, ctx.Param("id"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Location", fmt.Sprintf("api/v3/todos/%s", todo.Id))

	ctx.JSON(http.StatusCreated, todo)
}

func (c *Controller) GetList(ctx *gin.Context) {
	username := ctx.MustGet("username").(string)
	listId := ctx.Param("id")

	list, err := c.service.GetList(listId, username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (c *Controller) GetTodo(ctx *gin.Context) {
	userId := ctx.MustGet("username").(string)
	todoId := ctx.Param("id")

	todo, err := c.service.GetTodo(todoId, userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "todo not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) PatchList(ctx *gin.Context) {
	var reqBody data.TodoListPatchDTO

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	userId := ctx.MustGet("userId").(string)
	listId := ctx.Param("id")

	list, err := c.service.PatchList(&reqBody, listId, userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (c *Controller) PatchTodo(ctx *gin.Context) {
	var reqBody data.TodoPatchDTO

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	username := ctx.MustGet("username").(string)
	todoId := ctx.Param("id")

	todo, err := c.service.PatchTodo(&reqBody, todoId, username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "todo not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) DeleteList(ctx *gin.Context) {
	username := ctx.MustGet("username").(string)
	listId := ctx.Param("id")

	err := c.service.DeleteList(listId, username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) DeleteTodo(ctx *gin.Context) {
	username := ctx.MustGet("username").(string)
	todoId := ctx.Param("id")

	err := c.service.DeleteTodo(todoId, username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "todo not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) GetLists(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)
	query := ctx.DefaultQuery("completed", "")

	var completedBool *bool

	if query != "" {
		tempBool := query == "true"
		completedBool = &tempBool
	}

	lists, err := c.service.GetLists(userId, completedBool)
	if err != nil {
		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, lists)
}

func (c *Controller) GetTodos(ctx *gin.Context) {
	username := ctx.MustGet("username").(string)
	listId := ctx.Param("id")
	query := ctx.DefaultQuery("completed", "")

	var completedBool *bool

	if query != "" {

		switch query {
		case "true":
			tempBool := true
			completedBool = &tempBool
		case "false":
			tempBool := false
			completedBool = &tempBool
		default:
			ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid query "))
			return
		}

	}

	todos, err := c.service.GetTodos(listId, username, completedBool)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var reqBody data.UserCreateDTO

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	user, err := c.service.CreateUser(&reqBody)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *Controller) Login(ctx *gin.Context) {
	var reqBody data.UserLoginDTO

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	token, err := c.service.Login(&reqBody)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
