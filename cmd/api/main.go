package main

import (
	"github.com/aruncs31s/go_gin_api_starterkit/internal/application/container"
	"github.com/aruncs31s/go_gin_api_starterkit/internal/domain/models"
	"github.com/aruncs31s/go_gin_api_starterkit/internal/routes"
	"github.com/aruncs31s/gologger"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	// Get inmemory database connection
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	return db
}
func main() {
	engine := gin.Default()
	container := container.NewUserContainer(GetDB())
	userHandler := container.GetUserHandler()
	routes.NewRouter(engine, *userHandler).RegisterUserRoutes()

	gologger.Info("Server Starting ", zap.String("address", ":8080"))

	engine.Run(":8080")
}
