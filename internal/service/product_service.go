package service

import (
	"context"
	"fmt"
	"storeapi/internal/data/repository"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type productService struct {
	repo repository.ProductRepo
}

func NewProductService(repo repository.ProductRepo) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll(ctx context.Context) ([]*models.Product, error) {
	list, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	return list, nil
}

func (s *productService) GetById(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	model, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *productService) Create(ctx context.Context, model *models.Product) error {
	return s.repo.Create(ctx, model)
}

func (s *productService) Update(ctx context.Context, model *models.Product) error {
	return s.repo.Update(ctx, model)
}

func (s *productService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
