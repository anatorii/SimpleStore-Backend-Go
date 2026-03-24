package service

import (
	"context"
	"storeapi/internal/domain/models"
	"storeapi/internal/domain/repos"

	"github.com/google/uuid"
)

type supplierService struct {
	repo repos.SupplierRepoInt
}

func NewSupplierService(repo repos.SupplierRepoInt) SupplierService {
	return &supplierService{repo: repo}
}

func (s *supplierService) GetAll(ctx context.Context) ([]*models.Supplier, error) {
	return nil, nil
}

func (s *supplierService) GetById(ctx context.Context, id uuid.UUID) (*models.Supplier, error) {
	return nil, nil
}

func (s *supplierService) Create(ctx context.Context, model *models.Supplier) error {
	return nil
}

func (s *supplierService) Update(ctx context.Context, model *models.Supplier) error {
	return nil
}

func (s *supplierService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
