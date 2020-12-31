package repository

import "github.com/pzentenoe/onboarding-go-clean-architecture/app/domain/models"

type ProductRepository interface {
	FindProducts() ([]*models.Product, error)
}
