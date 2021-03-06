package main

import (
	config "lataltastore/config"
	routes "lataltastore/routes"
)

func main() {
	config.InitDb()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}

