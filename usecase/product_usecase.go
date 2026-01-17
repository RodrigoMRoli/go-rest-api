package usecase

import (
	"fmt"
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(productId int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(productId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return product, nil

}

func (pu *ProductUsecase) UpdateProduct(product *model.Product) (*model.Product, error) {

	updatedProduct, err := pu.repository.UpdateProduct(*product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return updatedProduct, nil
}

func (pu *ProductUsecase) DeleteProduct(productId int) error {
	err := pu.repository.DeleteProduct(productId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
