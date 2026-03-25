package repository

import (
	"context"
	"errors"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type supplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) SupplierRepo {
	return &supplierRepo{db: db}
}

func (r *supplierRepo) GetAll(ctx context.Context) ([]*models.Supplier, error) {
	var models []*models.Supplier
	query := `select id, name, address_id, phone_number
			  from suppliers`
	err := r.db.SelectContext(ctx, models, query)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (r *supplierRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Supplier, error) {
	var model models.Supplier
	query := `select id, name, address_id, phone_number
			  from suppliers
			  where id = $1`
	err := r.db.GetContext(ctx, &model, query, id)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *supplierRepo) Create(ctx context.Context, model *models.Supplier) error {
	query := `insert into suppliers (name, address_id, phone_number) values ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, model.Name, model.AddressId, model.PhoneNumber)
	return err
}

func (r *supplierRepo) Update(ctx context.Context, model *models.Supplier) error {
	query := `update suppliers set address_id = $1 where id = $2`
	_, err := r.db.ExecContext(ctx, query, model.AddressId, model.Id)
	return err
}

func (r *supplierRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from suppliers where id = $1`
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
