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
	var model models.Image
	query := `select id, image, description
			  from images
			  where id = $1`
	err := r.db.GetContext(ctx, &model, query, id)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *ImageRepo) GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error) {
	var model models.Image
	query := `select id, image, description
			  from images
			  where id = $1`
	err := r.db.GetContext(ctx, &model, query, productId)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *ImageRepo) Create(ctx context.Context, model *models.Image) error {
	query := `insert into images (image, description) values ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, model.Data, model.Description)
	return err
}

func (r *ImageRepo) Update(ctx context.Context, model *models.Image) error {
	query := `update images set image = $1, description = $2`
	_, err := r.db.ExecContext(ctx, query, model.Data, model.Description)
	return err
}

func (r *ImageRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from images where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
