package service

import (
	"context"
	"storeapi/internal/domain/models"
	"storeapi/internal/domain/repos"

	"github.com/google/uuid"
)

type productService struct {
	repo repos.ProductRepoInt
}

func NewProductService(repo repos.ProductRepoInt) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll(ctx context.Context) ([]*models.Product, error) {
	return nil, nil
}

func (s *productService) GetById(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return nil, nil
}

func (s *productService) Create(ctx context.Context, model *models.Product) error {
	return nil
}

func (s *productService) Update(ctx context.Context, model *models.Product) error {
	return nil
}

func (s *productService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
