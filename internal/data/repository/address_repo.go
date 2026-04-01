package repository

import (
	"context"
	"database/sql"
	"errors"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type addressRepo struct {
	db *sqlx.DB
}

func NewAddressRepo(db *sqlx.DB) AddressRepo {
	return &addressRepo{db: db}
}

func (r *addressRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Address, error) {
	var model models.Address
	query := `select id, country, city, street,
					 created_at, updated_at
			  from addresses
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

func (r *addressRepo) GetByAddress(ctx context.Context, country, city, street string) (*models.Address, error) {
	var model models.Address
	query := `select id, country, city, street,
					 created_at, updated_at
			  from addresses
			  where country = $1 and city = $2 and street = $3`
	err := r.db.Get(&model, query, country, city, street)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func (r *addressRepo) Create(ctx context.Context, model *models.Address) error {
	query := `insert into addresses (country, city, street)
			  values ($1, $2, $3)`
	_, err := r.db.Exec(query,
		model.Country, model.City, model.Street)
	return err
}

func (r *addressRepo) Update(ctx context.Context, model *models.Address) error {
	query := `update addresses set country = $1, city = $2, street = $3
			  where id = $4`
	_, err := r.db.Exec(query,
		model.Country, model.City, model.Street, model.Id)
	return err
}

func (r *addressRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from addresses where id = $1`
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
