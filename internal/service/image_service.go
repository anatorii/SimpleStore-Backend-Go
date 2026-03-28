package service

import (
	"context"
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
		return nil, err
	}
	return model, nil
}

func (s *imageService) GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error) {
	product, err := s.productRepo.GetById(ctx, productId)
	if err != nil {
		return nil, err
	}
	model, err := s.repo.GetById(ctx, product.ImageId)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *imageService) Create(ctx context.Context, model *models.Image) error {
	return s.repo.Create(ctx, model)
}

func (s *imageService) Update(ctx context.Context, model *models.Image) error {
	return s.repo.Update(ctx, model)
}

func (s *imageService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
