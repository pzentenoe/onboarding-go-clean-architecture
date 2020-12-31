package mapper

import (
	"github.com/pzentenoe/onboarding-project/app/domain/models"
	"github.com/pzentenoe/onboarding-project/app/infrastructure/database/sql/azure/entity"
)

func ProductsEntitiesTOProductsModel(productsEntities []*entity.Product) []*models.Product {
	products := make([]*models.Product, 0)

	for _, productsEntity := range productsEntities {

		products = append(products, productEntityToProductModel(productsEntity))
	}
	return products
}

func productEntityToProductModel(productsEntity *entity.Product) *models.Product {
	return &models.Product{
		ID:      productsEntity.IDProduct,
		Name:    productsEntity.NameProduct,
		Price:   productsEntity.PriceProduct,
		Barcode: productsEntity.SKU,
	}
}
