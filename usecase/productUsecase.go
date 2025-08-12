package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (int, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return 0, err
	}

	return productId, nil
}

func (pu *ProductUsecase) GetProductById(idProduct int) (*model.Product, error) {
	productId, err := pu.repository.GetProductById(idProduct)
	if err != nil {
		return nil, err
	}

	return productId, nil
}
