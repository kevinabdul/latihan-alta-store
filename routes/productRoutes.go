package routes

import(
	product "lataltastore/controllers/product"
)

func registerProductRoutes() {
	e.GET("/products", product.GetProductsController)

	e.GET("/products/:id", product.GetProductByIdController)
}

