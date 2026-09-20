package routes

import (
	"github.com/aruncs31s/go_gin_api_starterkit/internal/application/http/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine         *gin.Engine
	userController handler.UserHandler
}

func NewRouter(engine *gin.Engine, userController handler.UserHandler) *Router {
	return &Router{
		engine:         engine,
		userController: userController,
	}
}
