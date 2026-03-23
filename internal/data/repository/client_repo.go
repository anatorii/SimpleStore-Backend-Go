package repository

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) *ClientRepo {
	return &ClientRepo{db: db}
}

func (r *ClientRepo) GetAll(ctx context.Context) ([]*models.Client, error) {
	return nil, nil
}

func (r *ClientRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	return nil, nil
}

func (r *ClientRepo) GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error) {
	return nil, nil
}

func (r *ClientRepo) Create(ctx context.Context, model *models.Client) error {
	return nil
}

func (r *ClientRepo) Update(ctx context.Context, model *models.Client) error {
	return nil
}

func (r *ClientRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
