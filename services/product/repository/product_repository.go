package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/leifarriens/go-microservices/services/product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	err := db.AutoMigrate(&model.Product{})

	if err != nil {
		log.Fatalln(err)
	}

	return &productRepository{db: db}
}

func (r *productRepository) FindAll(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	result := r.db.Find(&products)

	amount := result.RowsAffected

	fmt.Printf("All: %d\n", amount)

	return products, nil
}
