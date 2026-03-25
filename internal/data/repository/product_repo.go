package repository

import (
	"context"
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
	err := r.db.SelectContext(ctx, models, query)
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
	err := r.db.GetContext(ctx, &model, query, id)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *productRepo) Create(ctx context.Context, model *models.Product) error {
	query := `insert into products (name, category, price,
									available_stock, last_update_date, supplier_id, image_id)
			  values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, model.Name, model.Category, model.Price,
		model.AvailableStock, model.LastUpdateDate, model.SupplierId, model.ImageId)
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
				image_id = $7,
			  where id = $8`
	_, err := r.db.ExecContext(ctx, query,
		model.Name,
		model.Category,
		model.Price,
		model.AvailableStock,
		model.LastUpdateDate,
		model.SupplierId,
		model.ImageId,
		model.Id)
	return err
}

func (r *productRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from products where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
