package app

import (
	"AnotherTest/infra/product"
	"fmt"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	registerProductAPI(e, "gorm")
	e.Logger.Fatal(e.Start(":1323"))
}

// IF orm == bun then init bun repo
// IF orm == gorm then init gorm repo
func registerProductAPI(e *echo.Echo, orm string) {
	var productRepository product.ProductRepository
	fmt.Println("masuk proses")
	switch orm {
	case "bun":
		productRepository = product.NewProductBunRepositoryImpl("postgres://postgres:12345678@localhost:5432/testbun?sslmode=disable")
		break
	case "gorm":
		productRepository = product.NewProductGormRepositoryImpl("host=localhost user=postgres password=12345678 dbname=testbun port=5432 sslmode=disable TimeZone=Asia/Jakarta")
		break
	default:
		panic(`unknown orm selections!`)
	}

	productService := product.NewProductServiceImpl(productRepository) //bun
	productApi := product.NewProductAPIImpl(productService)
	product.RegisterRoute(e, productApi)
}
