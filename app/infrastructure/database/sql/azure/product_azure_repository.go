package azure

import (
	"fmt"
	"github.com/pzentenoe/onboarding-project/app/domain/models"
	"github.com/pzentenoe/onboarding-project/app/infrastructure/database/sql/azure/entity"
	"github.com/pzentenoe/onboarding-project/app/infrastructure/database/sql/azure/mapper"
)

type productAzureRepository struct {
}

func NewProductAzureRepository() *productAzureRepository {
	return &productAzureRepository{}
}

func (r *productAzureRepository) FindProducts() ([]*models.Product, error) {

	productsEntities, err := getProductsDB()
	if err != nil {
		return nil, err
	}

	products := mapper.ProductsEntitiesTOProductsModel(productsEntities)
	return products, nil

}



func getProductsDB() ([]*entity.Product, error) {

	productsEntities := make([]*entity.Product, 0)
	for i := 0; i < 10; i++ {
		productsEntities = append(productsEntities, &entity.Product{
			IDProduct:    fmt.Sprintf("%d%d%d", i+100, i+200, i+300),
			NameProduct:  fmt.Sprintf("product-%d", i+1),
			PriceProduct: (i + 1) * 100,
			SKU:          fmt.Sprintf("%d%d%d", i+20, i+30, i+60),
		})
	}
	return productsEntities, nil

}
