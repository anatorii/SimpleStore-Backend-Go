package repository

import (
	"context"
	"database/sql"
	"errors"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) GetAll(ctx context.Context) ([]*models.Product, error) {
	var models []*models.Product
	query := `select id, name, category, price,
					 available_stock, last_update_date, supplier_id, image_id
			  from products`
	err := r.db.Select(&models, query)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (r *productRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	var model models.Product
	query := `select id, name, category, price,
					 available_stock, last_update_date, supplier_id, image_id
			  from products
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

func (r *productRepo) Create(ctx context.Context, model *models.Product) error {
	query := `insert into products (name, category, price,
									available_stock, last_update_date, supplier_id, image_id)
			  values ($1, $2, $3, $4, $5, $6, $7)`

	SupplierId := uuid.NullUUID{Valid: false}
	if model.SupplierId != uuid.Nil {
		SupplierId = uuid.NullUUID{UUID: model.SupplierId, Valid: true}
	}

	ImageId := uuid.NullUUID{Valid: false}
	if model.ImageId != uuid.Nil {
		ImageId = uuid.NullUUID{UUID: model.ImageId, Valid: true}
	}

	_, err := r.db.Exec(query, model.Name, model.Category, model.Price,
		model.AvailableStock, model.LastUpdateDate, SupplierId, ImageId)
	return err
}

func (r *productRepo) Update(ctx context.Context, model *models.Product) error {
	query := `update products set
				name = $1,
				category = $2,
				price = $3,
				available_stock = $4,
				last_update_date = $5,
				supplier_id = $6,
				image_id = $7
			  where id = $8`

	SupplierId := uuid.NullUUID{Valid: false}
	if model.SupplierId != uuid.Nil {
		SupplierId = uuid.NullUUID{UUID: model.SupplierId, Valid: true}
	}

	ImageId := uuid.NullUUID{Valid: false}
	if model.ImageId != uuid.Nil {
		ImageId = uuid.NullUUID{UUID: model.ImageId, Valid: true}
	}

	result, err := r.db.Exec(query,
		model.Name,
		model.Category,
		model.Price,
		model.AvailableStock,
		model.LastUpdateDate,
		SupplierId,
		ImageId,
		model.Id)
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

func (r *productRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from products where id = $1`
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
