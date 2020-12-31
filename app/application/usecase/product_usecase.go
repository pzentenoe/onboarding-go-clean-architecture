package usecase

import "github.com/pzentenoe/onboarding-project/app/domain/models"

type ProductUsecase interface {
	GetProducts() ([]*models.Product, error)
}
