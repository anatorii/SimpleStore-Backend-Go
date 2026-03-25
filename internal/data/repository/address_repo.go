package repository

import (
	"context"
	"errors"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AddressRepo struct {
	db *sqlx.DB
}

func NewAddressRepo(db *sqlx.DB) *AddressRepo {
	return &AddressRepo{db: db}
}

func (r *AddressRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Address, error) {
	var model models.Address
	query := `select id, coutry, city, street
			  from addresses where id = $1`
	err := r.db.GetContext(ctx, &model, query, id)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *AddressRepo) Create(ctx context.Context, model *models.Address) error {
	query := `insert into addresses (coutry, city, street)
			  values ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query,
		model.Country, model.City, model.Street)
	return err
}

func (r *AddressRepo) Update(ctx context.Context, model *models.Address) error {
	query := `update addresses set coutry = $1, city = $2, street = $3
			  where id = $4`
	_, err := r.db.ExecContext(ctx, query,
		model.Country, model.City, model.Street, model.Id)
	return err
}

func (r *AddressRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from addresses where id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
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
