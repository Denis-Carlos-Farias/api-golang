package controllers

import (
	"net/http"

	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/models"
	"github.com/gin-gonic/gin"
)

type produtController struct {
	products []models.Product
}

func NewprodutController() *produtController {
	return &produtController{}
}

func (pc *produtController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, pc.products)
}

func (pc *produtController) Create(ctx *gin.Context) {

	product := models.NewProduct()

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	pc.products = append(pc.products, *product)

	ctx.JSON(http.StatusOK, product)
}
