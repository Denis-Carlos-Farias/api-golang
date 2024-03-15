package main

import (
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.Default()

	routes.AppRoutes(g)

	g.Run(":3000")
}
