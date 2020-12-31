package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-go-clean-architecture/app/application/usecase/product"
	"github.com/pzentenoe/onboarding-go-clean-architecture/app/infrastructure/database/sql/azure"
	"github.com/pzentenoe/onboarding-go-clean-architecture/app/infrastructure/http/handlers"
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
