package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leifarriens/go-microservices/internal/shared"
	_ "github.com/leifarriens/go-microservices/services/product/docs"
	"github.com/leifarriens/go-microservices/services/product/handler"
	"github.com/leifarriens/go-microservices/services/product/repository"
	"github.com/leifarriens/go-microservices/services/product/service"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

//	@title			Product Service
//	@version		2.0
//	@description	This is the product service

// @host	localhost:1323
func main() {
	publicKey := shared.LoadPublicKey()

	connStr := shared.GetConnectionString()

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	productRepository := repository.NewProductRepository(db)

	productService := service.NewProductService(&service.ProductServiceConfig{
		ProductRepository: productRepository,
	})

	s := shared.Server(&shared.ServerConfig{
		Validator: true,
		Swagger:   true,
		CORSConfig: &middleware.CORSConfig{
			// https://echo.labstack.com/docs/cookbook/cors#server-using-a-custom-function-to-allow-origins
			AllowOrigins: []string{"http://localhost:5173"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		},
	})

	handler.NewHandler(&handler.HandlerConfig{
		E:              s,
		ProductService: productService,
		PublicKey:      publicKey,
	})

	s.Logger.Fatal(s.Start(":1323"))
}

// func seed(db *gorm.DB) {
// 	product := model.Product{
// 		Name:      "Product 1",
// 		Price:     100,
// 		Available: true,
// 	}

// 	db.Create(&product)
// }
