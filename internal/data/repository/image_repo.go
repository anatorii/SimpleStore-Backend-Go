package repository

import (
	"context"
	"database/sql"
	"errors"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type imageRepo struct {
	db *sqlx.DB
}

func NewImageRepo(db *sqlx.DB) ImageRepo {
	return &imageRepo{db: db}
}

func (r *imageRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	var model models.Image
	query := `select id, data, coalesce(description, '') description
			  from images
			  where id = $1`
	err := r.db.Get(&model, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func (r *imageRepo) GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error) {
	var model models.Image
	query := `select id, data, coalesce(description, '') description
			  from images
			  where id = $1`
	err := r.db.Get(&model, query, productId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func (r *imageRepo) Create(ctx context.Context, model *models.Image) (uuid.UUID, error) {
	id := uuid.New()
	query := `insert into images (id, data) values ($1, $2)`
	_, err := r.db.Exec(query, id, model.Data)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *imageRepo) Update(ctx context.Context, model *models.Image) error {
	query := `update images set data = $1, description = $2`
	result, err := r.db.Exec(query, model.Data, model.Description)
	if err != nil {
		return err
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("NO_AFFECTED")
	}
	return nil
}

func (r *imageRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from images where id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("NO_AFFECTED")
	}
	return nil
}
