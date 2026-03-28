package service

import (
	"context"
	"fmt"
	"storeapi/internal/data/repository"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type supplierService struct {
	repo     repository.SupplierRepo
	addrRepo repository.AddressRepo
}

func NewSupplierService(repo repository.SupplierRepo, addrRepo repository.AddressRepo) SupplierService {
	return &supplierService{repo: repo, addrRepo: addrRepo}
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

func (s *supplierService) UpdateAddress(ctx context.Context, model *models.Supplier, address models.Address) error {
	var addr *models.Address
	addr, err := s.addrRepo.GetByAddress(ctx, address.Country, address.City, address.Street)
	if err != nil {
		return fmt.Errorf("failed to update supplier address: %w", err)
	}
	if addr == nil {
		err = s.addrRepo.Create(ctx, &address)
		if err != nil {
			return fmt.Errorf("failed to update supplier address: %w", err)
		}
		addr, err = s.addrRepo.GetByAddress(ctx, address.Country, address.City, address.Street)
		if err != nil {
			return fmt.Errorf("failed to update supplier address: %w", err)
		}
	}
	model.AddressId = addr.Id
	return s.Update(ctx, model)
}

func (s *supplierService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
