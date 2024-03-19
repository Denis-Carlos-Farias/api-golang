package main

import (
	_ "github.com/Denis-Carlos-Farias/crud-golang/cmd/api/docs"

	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/controllers"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/routes"
)

// @title 	Product API
// @version	1.0
// @description A Product service API in Go using Gin framework
// @host 	localhost:3000
func main() {

	controller := controllers.NewprodutController()

	gRouter := routes.AppRoutes(controller)

	gRouter.Run(":3000")
}
