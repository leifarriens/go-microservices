package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/leifarriens/go-microservices/internal/shared"
	_ "github.com/leifarriens/go-microservices/services/product/docs"
	"github.com/leifarriens/go-microservices/services/product/handler"
	"github.com/leifarriens/go-microservices/services/product/repository"
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

	s := shared.Server(&shared.ServerConfig{
		Validator: true,
		KeyAuth:   false,
		Swagger:   true,
		AllowOrigins: []string{
			"http://localhost:5173",
		},
	})

	handler.NewHandler(&handler.HandlerConfig{
		E:                 s,
		ProductRepository: productRepository,
		PublicKey:         publicKey,
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
