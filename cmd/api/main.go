package main

import (
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// connectionString := "postgresql://postgres:123456@localhost:5432/postgres"

	g := gin.Default()

	routes.AppRoutes(g)

	g.Run(":3000")
}
