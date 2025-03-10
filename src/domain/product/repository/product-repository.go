package repository

import (
	"test-kp-golang/src/domain/product/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts() ([]entity.Product, error) {
	var products []entity.Product

	result := r.db.Find(&products)
	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (entity.Product, error) {
	var product entity.Product

	result := r.db.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}
