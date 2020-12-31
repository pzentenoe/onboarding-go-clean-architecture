package repository

import "github.com/pzentenoe/onboarding-project/app/domain/models"

type ProductRepository interface {
	FindProducts() ([]*models.Product, error)
}
