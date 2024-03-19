package routes

import (
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AppRoutes(product *controllers.ProdutController) *gin.Engine {

	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")
	v1 := baseRouter.Group("/v1")
	v1.GET("/product/:id", product.GetById)
	v1.GET("/products", product.GetAll)
	v1.POST("/product", product.Create)
	v1.DELETE("/product/:id", product.Delete)
	v1.PUT("/product", product.Update)

	return router
}
