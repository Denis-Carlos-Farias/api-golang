package repository

import (
	"errors"

	"github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/entities"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/domain/interfaces/repositories"
	"github.com/Denis-Carlos-Farias/crud-golang/cmd/helper"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func NewProductRepository(Db *gorm.DB) repositories.IProductRepository {
	return &ProductRepository{Db: Db}
}

// Delete implements repositories.IProductRepository.
func (p *ProductRepository) Delete(productId int) {
	var product entities.Product
	result := p.Db.Where("id = ?", productId).Delete(&product)
	helper.ErrorPanic(result.Error)
}

// FindAll implements repositories.IProductRepository.
func (p *ProductRepository) FindAll() []entities.Product {
	var products []entities.Product
	result := p.Db.Find(&products)
	helper.ErrorPanic(result.Error)
	return products
}

// FindById implements repositories.IProductRepository.
func (p *ProductRepository) FindById(productId int) (product entities.Product, err error) {
	var pdt entities.Product
	result := p.Db.Find(&pdt, productId)
	if result != nil {
		return pdt, nil
	} else {
		return pdt, errors.New("tag is not found")
	}
}

// Save implements repositories.IProductRepository.
func (p *ProductRepository) Save(product entities.Product) {
	result := p.Db.Create(&product)
	helper.ErrorPanic(result.Error)
}

// Update implements repositories.IProductRepository.
func (p *ProductRepository) Update(product entities.Product) {
	var updateProduct = entities.Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	result := p.Db.Model(&product).Updates(updateProduct)
	helper.ErrorPanic(result.Error)
}
