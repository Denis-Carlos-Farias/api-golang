package routes

import (
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {

	product := controllers.NewprodutController()

	v1 := router.Group("/v1")
	{
		v1.GET("/product/:id", product.GetById)
		v1.GET("/products", product.GetAll)
		v1.POST("/product", product.Create)
		v1.DELETE("/product/:id", product.Delete)
		v1.PUT("/product", product.Update)
	}

	return v1
}
