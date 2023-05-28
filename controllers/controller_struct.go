package controllers

import "github.com/mateicheles1/golang-crud/service"

type Controller struct {
	service service.TodoListService
}

func NewController(service service.TodoListService) *Controller {
	return &Controller{
		service: service,
	}
}
