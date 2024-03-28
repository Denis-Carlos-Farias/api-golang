package services

import "github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/entities"

type IProductService interface {
	Create(product entities.Product)
	Update(product entities.Product)
	Delete(productId int)
	FindById(productId int) entities.Product
	FindAll() []entities.Product
}
