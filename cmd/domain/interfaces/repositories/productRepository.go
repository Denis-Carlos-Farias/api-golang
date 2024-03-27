package repositories

import "github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/entities"

type IProductRepository interface {
	Save(product entities.Product)
	Update(product entities.Product)
	Delete(productId int)
	FindById(productId int) (product entities.Product, err error)
	FindAll() []entities.Product
}
