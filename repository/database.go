package repository

import (
	"github.com/mateicheles1/golang-crud/config"
	"github.com/mateicheles1/golang-crud/logs"
	"github.com/mateicheles1/golang-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Database.ConnectionString()), &gorm.Config{})

	if err != nil {
		logs.Logger.Fatal().Msgf("Failed to connect to DB: %s", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.JWT{}, &models.TodoList{}, &models.Todo{}); err != nil {
		logs.Logger.Fatal().Msgf("Failed to migrate schema: %s", err)
	}

	return db
}
