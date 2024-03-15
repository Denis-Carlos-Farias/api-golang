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

func (pc *produtController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, pc.products)
}

func (pc *produtController) GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	for _, product := range pc.products {
		if product.ID == id {
			ctx.JSON(http.StatusOK, product)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, "Product not found.")
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

func (pc *produtController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	for i, product := range pc.products {
		if product.ID == id {
			pc.products = append(pc.products[0:i], pc.products[i+1:]...)
			ctx.JSON(http.StatusOK, pc.products)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, "Product not found.")
}

func (pc *produtController) Update(ctx *gin.Context) {

	var pdt models.Product

	err := ctx.BindJSON(&pdt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	for i, product := range pc.products {
		if product.ID == pdt.ID {
			pc.products[i] = pdt
			ctx.JSON(http.StatusOK, pc.products[i])
			return
		}
	}

	ctx.JSON(http.StatusNotFound, "Product not found.")
}
