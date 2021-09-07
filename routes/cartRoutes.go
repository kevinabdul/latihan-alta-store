package routes

import(
	cart "lataltastore/controllers/cart"
	"lataltastore/middlewares"
)

func registerCartRoutes() {
	// e.GET("/carts", cart.GetCartsController, middlewares.AuthenticateUser)

	e.GET("/carts/:id", cart.GetCartByUserIdController, middlewares.AuthenticateUser, middlewares.CheckId)

	// e.POST("/carts", cart.CreateCartController, middlewares.AuthenticateUser)

	e.PUT("/carts/:id", cart.UpdateCartByUserIdController, middlewares.AuthenticateUser, middlewares.CheckId)
}

