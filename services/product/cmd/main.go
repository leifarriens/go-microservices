package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

//	@title			Product Service
//	@version		2.0
//	@description	This is the product service

// @host	localhost:1324
func main() {
	publicKey := shared.LoadPublicKey()

	connStr := shared.GetDBConnectionString()

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

	s.GET("/ping", func(c echo.Context) error {
		var dbTime time.Time

		db.Raw("SELECT NOW()").Scan(&dbTime)

		return c.JSON(http.StatusOK, fmt.Sprintf("OK %s", dbTime))
	})

	s.Logger.Fatal(s.Start(":1324"))
}

// func seed(db *gorm.DB) {
// 	product := model.Product{
// 		Name:      "Product 1",
// 		Price:     100,
// 		Available: true,
// 	}

// 	db.Create(&product)
// }
