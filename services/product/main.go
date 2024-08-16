package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/leifarriens/go-microservices/internal/shared"
	"github.com/leifarriens/go-microservices/services/product/handler"
	"github.com/leifarriens/go-microservices/services/product/repository"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	connStr := shared.GetConnectionString()

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	productRepository := repository.NewProductRepository(db)

	s := shared.Server(&shared.ServerConfig{
		KeyAuth: true,
	})

	handler.NewHandler(&handler.HandlerConfig{
		E:                 s,
		ProductRepository: productRepository,
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
