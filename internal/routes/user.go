package routes

func (r *Router) RegisterUserRoutes() {
	userGroup := r.engine.Group("/users")
	{
		userGroup.POST("/", r.userController.CreateUser)
		userGroup.GET("/:id", r.userController.GetUserByID)
		userGroup.PUT("/:id", r.userController.UpdateUser)
		userGroup.DELETE("/:id", r.userController.DeleteUser)
		userGroup.GET("/", r.userController.GetAllUsers)
	}
}
