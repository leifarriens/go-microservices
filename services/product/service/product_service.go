package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/leifarriens/go-microservices/services/product/model"
	"github.com/leifarriens/go-microservices/services/product/repository"
	"gorm.io/gorm"
)

var ErrProductNotFound = errors.New("product not found")

type ProductService interface {
	Add(ctx context.Context, product *model.ProductDto) (*model.Product, error)
	Get(ctx context.Context) ([]*model.Product, error)
	GetById(ctx context.Context, id string) (*model.Product, error)
}

type productService struct {
	ProductRepository repository.ProductRepository
}

type ProductServiceConfig struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(config *ProductServiceConfig) ProductService {
	return &productService{
		ProductRepository: config.ProductRepository,
	}
}

func (s *productService) Add(ctx context.Context, p *model.ProductDto) (*model.Product, error) {
	id, err := s.ProductRepository.Create(ctx, &model.Product{
		Name:      p.Name,
		Price:     p.Price,
		Available: p.Available,
	})

	if err != nil {
		return nil, err
	}

	product, err := s.ProductRepository.FindById(ctx, strconv.FormatUint(uint64(*id), 10))

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) Get(ctx context.Context) ([]*model.Product, error) {
	products, err := s.ProductRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) GetById(ctx context.Context, id string) (*model.Product, error) {
	product, err := s.ProductRepository.FindById(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	return product, nil
}
