package routes

import(
	cart "lataltastore/controllers/cart"
	"lataltastore/middlewares"
)

func registerCartRoutes() {
	e.GET("/cart", cart.GetCartController, middlewares.AuthenticateUser)

	e.POST("/cart", cart.CreateCartController, middlewares.AuthenticateUser)
}

