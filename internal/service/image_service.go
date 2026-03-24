package service

import (
	"context"
	"storeapi/internal/domain/models"
	"storeapi/internal/domain/repos"

	"github.com/google/uuid"
)

type imageService struct {
	repo repos.ImageRepoInt
}

func NewImageService(repo repos.ImageRepoInt) ImageService {
	return &imageService{repo: repo}
}

func (s *imageService) GetById(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	return nil, nil
}

func (s *imageService) GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error) {
	return nil, nil
}

func (s *imageService) Create(ctx context.Context, model *models.Image) error {
	return nil
}

func (s *imageService) Update(ctx context.Context, model *models.Image) error {
	return nil
}

func (s *imageService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
