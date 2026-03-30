package service

import (
	"context"
	"fmt"
	"storeapi/internal/data/repository"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type imageService struct {
	repo        repository.ImageRepo
	productRepo repository.ProductRepo
}

func NewImageService(repo repository.ImageRepo, productRepo repository.ProductRepo) ImageService {
	return &imageService{repo: repo, productRepo: productRepo}
}

func (s *imageService) GetById(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	model, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	return model, nil
}

func (s *imageService) GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error) {
	product, err := s.productRepo.GetById(ctx, productId)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	model, err := s.repo.GetById(ctx, product.ImageId)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	return model, nil
}

func (s *imageService) Create(ctx context.Context, model *models.Image, product *models.Product) error {
	id, err := s.repo.Create(ctx, model)
	if err != nil {
		return err
	}
	product, err = s.productRepo.GetById(ctx, product.Id)
	if err != nil {
		return fmt.Errorf("failed to create image: %w", err)
	}
	product.ImageId = id
	err = s.productRepo.Update(ctx, product)
	if err != nil {
		return fmt.Errorf("failed to create image: %w", err)
	}
	return nil
}

func (s *imageService) Update(ctx context.Context, model *models.Image) error {
	err := s.repo.Update(ctx, model)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			return err
		}
		return fmt.Errorf("failed to update image: %w", err)
	}
	return nil
}

func (s *imageService) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			return err
		}
		return fmt.Errorf("failed to delete image: %w", err)
	}
	return nil
}
