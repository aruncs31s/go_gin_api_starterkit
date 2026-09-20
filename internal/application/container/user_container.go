package container

import (
	"github.com/aruncs31s/go_gin_api_starterkit/internal/application/http/handler"
	"github.com/aruncs31s/go_gin_api_starterkit/internal/repository"
	"github.com/aruncs31s/go_gin_api_starterkit/internal/service"
	"gorm.io/gorm"
)

type UserContainer struct {
	db          *gorm.DB
	userService service.UserService
}

func NewUserContainer(db *gorm.DB) *UserContainer {
	userService := service.NewUserService(
		repository.NewUserRepository(db),
	)
	return &UserContainer{
		db:          db,
		userService: userService,
	}
}

func (u *UserContainer) GetUserHandler() *handler.UserHandler {
	return handler.NewUserHandler(u.userService)
}

func (u *UserContainer) GetDB() *gorm.DB {
	return u.db
}
