package service

import (
	"context"
	"fmt"
	"storeapi/internal/data/repository"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type supplierService struct {
	repo repository.SupplierRepo
}

func NewSupplierService(repo repository.SupplierRepo) SupplierService {
	return &supplierService{repo: repo}
}

func (s *supplierService) GetAll(ctx context.Context) ([]*models.Supplier, error) {
	list, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get suppliers: %w", err)
	}
	return list, nil
}

func (s *supplierService) GetById(ctx context.Context, id uuid.UUID) (*models.Supplier, error) {
	model, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *supplierService) Create(ctx context.Context, model *models.Supplier) error {
	return s.repo.Create(ctx, model)
}

func (s *supplierService) Update(ctx context.Context, model *models.Supplier) error {
	return s.repo.Update(ctx, model)
}

func (s *supplierService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
