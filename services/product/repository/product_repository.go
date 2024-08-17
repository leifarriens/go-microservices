package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/leifarriens/go-microservices/services/product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) (*uint, error)
	FindAll(ctx context.Context) ([]*model.Product, error)
	FindById(ctx context.Context, id string) (*model.Product, error)
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

func (r *productRepository) Create(ctx context.Context, product *model.Product) (*uint, error) {
	result := r.db.Create(&product)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return &product.ID, nil
}

func (r *productRepository) FindAll(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	result := r.db.Find(&products)

	amount := result.RowsAffected

	fmt.Printf("All: %d\n", amount)

	return products, nil
}

func (r *productRepository) FindById(ctx context.Context, id string) (*model.Product, error) {
	var products *model.Product

	r.db.First(&products, id)

	return products, nil
}
