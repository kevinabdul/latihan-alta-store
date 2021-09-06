package routes

import(
	cart "lataltastore/controllers/cart"
	"lataltastore/middlewares"
)

func registerCartRoutes() {
	e.GET("/carts", cart.GetCartsController, middlewares.AuthenticateUser)

	e.GET("/carts/:id", cart.GetCartByIdController, middlewares.AuthenticateUser, middlewares.CheckId)

	e.POST("/carts", cart.CreateCartController, middlewares.AuthenticateUser)
}

