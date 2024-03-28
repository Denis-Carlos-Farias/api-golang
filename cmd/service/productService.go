package service

import (
	"fmt"

	"github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/entities"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/interfaces/repositories"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/interfaces/services"
)

type ProductService struct {
	Repository repositories.IProductRepository
}

func NewProductSrvice(Repository repositories.IProductRepository) services.IProductService {
	return &ProductService{Repository: Repository}
}

// Create implements services.IProductService.
func (p *ProductService) Create(product entities.Product) {
	p.Repository.Save(product)
}

// Update implements services.IProductService.
func (p *ProductService) Update(product entities.Product) {
	p.Repository.Update(product)
}

// Delete implements services.IProductService.
func (p *ProductService) Delete(productId int) {
	p.Repository.Delete(productId)
}

// FindAll implements services.IProductService.
func (p *ProductService) FindAll() []entities.Product {
	result := p.Repository.FindAll()

	return result
}

// FindById implements services.IProductService.
func (p *ProductService) FindById(productId int) entities.Product {
	result, err := p.Repository.FindById(productId)
	if err != nil {
		fmt.Printf("FindById method got an error: %s", err.Error())
	}

	return result
}
