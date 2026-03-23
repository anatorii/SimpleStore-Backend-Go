package repository

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ImageRepo struct {
	db *sqlx.DB
}

func NewImageRepo(db *sqlx.DB) *ImageRepo {
	return &ImageRepo{db: db}
}

func (r *ImageRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	return nil, nil
}

func (r *ImageRepo) GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error) {
	return nil, nil
}

func (r *ImageRepo) Create(ctx context.Context, model *models.Image) error {
	return nil
}

func (r *ImageRepo) Update(ctx context.Context, model *models.Image) error {
	return nil
}

func (r *ImageRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
