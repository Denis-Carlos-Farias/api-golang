package controllers

import (
	"net/http"

	"github.com/Denis-Carlos-Farias/crud-golang/cmd/api/models"
	"github.com/gin-gonic/gin"
)

type ProdutController struct {
	products []models.Product
}

func NewprodutController() *ProdutController {
	return &ProdutController{}
}

// FindAllProducts 	godoc
// @Summary			Get All products.
// @Description		Return list of products.
// @Tags			product
// @Success			200 {object} models.Product{}
// @Router			/api/v1/products [get]
func (pc *ProdutController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, pc.products)
}

// FindProductById 		godoc
// @Summary				Get Single product by id.
// @Param				id path string true "update product by id"
// @Description			Return the tahs whoes productId valu mathes id.
// @Produce				application/json
// @Tags				product
// @Success			200 {object} models.Product{}
// @Router				/api/v1/product/{id} [get]
func (pc *ProdutController) GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	for _, product := range pc.products {
		if product.ID == id {
			ctx.JSON(http.StatusOK, product)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, "Product not found.")
}

// CreateProduct	godoc
// @Summary			Create Product
// @Description		Save product data in mmory.
// @Param			tags body models.Product true "Create product"
// @Produce			application/json
// @Tags			product
// @Success			200 {object} models.Product{}
// @Router			/api/v1/product [post]
func (pc *ProdutController) Create(ctx *gin.Context) {

	product := models.NewProduct()

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	pc.products = append(pc.products, *product)

	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct	godoc
// @Summary			Delete Product
// @Description		Remove Product data by id.
// @Produce			application/json
// @Tags			product
// @Success			200 {object} models.Product{}
// @Router			/api/v1/Product/{id} [delete]
func (pc *ProdutController) Delete(ctx *gin.Context) {

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

// UpdateProduct	godoc
// @Summary			Update Product
// @Description		Update product data in mmory.
// @Param			tags body models.Product true "Update product"
// @Produce			application/json
// @Tags			product
// @Success			200 {object} models.Product{}
// @Router			/api/v1/product [put]
func (pc *ProdutController) Update(ctx *gin.Context) {

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
