package usecase

import (
	"test-kp-golang/src/domain/product/repository"
	"test-kp-golang/src/domain/product/response"
)

type ProductUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
	}
}

func (u *ProductUsecase) GetProducts() ([]response.Product, error) {
	products, err := u.productRepo.GetProducts()
	if err != nil {
		return []response.Product{}, err
	}

	produtctResponse := []response.Product{}
	for _, product := range products {
		produtctResponse = append(produtctResponse, response.Product{
			ID:          product.ID,
			Name:        product.Name,
			AmountPrice: product.AmountPrice,
			AmountShip:  product.AmountShip,
		})
	}

	return produtctResponse, nil
}

func (u *ProductUsecase) GetProductByID(id int) (response.Product, error) {
	product, err := u.productRepo.GetProductByID(id)
	if err != nil {
		return response.Product{}, err
	}

	produtctResponse := response.Product{
		ID:          product.ID,
		Name:        product.Name,
		AmountPrice: product.AmountPrice,
		AmountShip:  product.AmountShip,
	}

	return produtctResponse, nil
}
