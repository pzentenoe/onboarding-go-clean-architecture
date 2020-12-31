package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-project/app/application/usecase/product"
	"github.com/pzentenoe/onboarding-project/app/infrastructure/database/sql/azure"
	"github.com/pzentenoe/onboarding-project/app/infrastructure/http/handlers"
)

func main() {

	e := echo.New()
	e.HideBanner = true

	handlers.NewHealthHandler(e)

	productRepository := azure.NewProductAzureRepository()
	productUsecase := product.NewProductUsecase(productRepository)
	handlers.NewProductHandler(e, productUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
