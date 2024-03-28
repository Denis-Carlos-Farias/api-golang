package main

import (
	"fmt"

	_ "github.com/Denis-Carlos-Farias/crud-golang/cmd/api/docs"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/repository"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/service"

	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/controllers"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/routes"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/entities"
	"github.com/Denis-Carlos-Farias/crud-golang/config"
)

// @title 	Product API
// @version	1.0
// @description A Product service API in Go using Gin framework
// @host 	localhost:3000
func main() {

	// Configuração da conexão com o banco de dados
	db := config.DatabaseConnection()
	fmt.Println("Conexão bem-sucedida ao banco de dados SQL Server!")
	db.Table("Product").AutoMigrate(&entities.Product{})

	repository := repository.NewProductRepository(db)
	service := service.NewProductSrvice(repository)
	controller := controllers.NewprodutController(service)

	gRouter := routes.AppRoutes(controller)

	gRouter.Run(":3000")
}
