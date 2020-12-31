package usecase

import "github.com/pzentenoe/onboarding-go-clean-architecture/app/domain/models"

type ProductUsecase interface {
	GetProducts() ([]*models.Product, error)
}
