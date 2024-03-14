package routes

import (
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {

	product := controllers.NewprodutController()

	v1 := router.Group("/v1")
	{
		v1.GET("/", product.FindAll)
		v1.POST("/", product.Create)
	}

	return v1
}
