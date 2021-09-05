package routes

import(
	user "lataltastore/controllers/user"
	"lataltastore/middlewares"
)

func registerUserRoutes() {
	e.GET("/users", user.GetUsersController, middlewares.AuthenticateUser)

	e.POST("/users", user.AddUserController)

	e.POST("/login", user.LoginUserController)

	r := e.Group("/users/:id")

	r.Use(middlewares.AuthenticateUser)

	r.Use(middlewares.CheckId)

	r.GET("", user.GetUserByIdController)

	r.PUT("", user.EditUserController)

	r.DELETE("", user.DeleteUserController)
	
}

