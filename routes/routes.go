package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mateicheles1/golang-crud/config"
	"github.com/mateicheles1/golang-crud/controllers"
	"github.com/mateicheles1/golang-crud/logs"
	"github.com/mateicheles1/golang-crud/middleware"
	"github.com/mateicheles1/golang-crud/repository"
	"github.com/mateicheles1/golang-crud/service"
)

func SetupRoutes() {

	config := config.NewConfig("./config/config.json")
	db := repository.NewDatabase(config)
	data := repository.NewTodoListDBImpl(db)
	service := service.NewTodolistServiceImpl(data)
	controller := controllers.NewController(service)

	router := gin.New()

	router.Use(middleware.InfoHandler())
	router.Use(middleware.ErrorHandler())
	router.Use(gin.Recovery())

	router.POST("api/v3/users/signup", controller.CreateUser)
	router.POST("api/v3/users/login", controller.Login)

	appRouter := router.Group("api/v3")

	appRouter.Use(middleware.AuthMiddleware())

	appRouter.GET("lists", controller.GetLists)
	appRouter.GET("lists/:id/todos", controller.GetTodos)

	appRouter.GET("lists/:id", controller.GetList)
	appRouter.POST("lists", controller.CreateList)
	appRouter.PATCH("lists/:id", controller.PatchList)
	appRouter.DELETE("lists/:id", controller.DeleteList)

	appRouter.GET("todos/:id", controller.GetTodo)
	appRouter.POST("lists/:id/todos", controller.CreateTodo)
	appRouter.PATCH("todos/:id", controller.PatchTodo)
	appRouter.DELETE("todos/:id", controller.DeleteTodo)

	if err := router.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)); err != nil {
		logs.Logger.Fatal().Msgf("Failed to start server due to: %s", err)
	}

}
