package product

import (
	"github.com/pzentenoe/onboarding-go-clean-architecture/app/domain/models"
	"github.com/pzentenoe/onboarding-go-clean-architecture/app/domain/repository"
)

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) *productUsecase {
	u := &productUsecase{productRepository: productRepository}
	return u
}

func (u *productUsecase) GetProducts() ([]*models.Product, error) {
	return u.productRepository.FindProducts()
}
