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
	var models []*models.Client
	query := `select id, client_name, client_surname, birthday, 
			         gender, registration_date, address_id
			  from clients`
	err := r.db.SelectContext(ctx, models, query)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (r *ClientRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	var model models.Client
	query := `select id, client_name, client_surname, birthday,
	                 gender, registration_date, address_id
			  from clients
			  where id = $1`
	err := r.db.GetContext(ctx, &model, query, id)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *ClientRepo) GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error) {
	var model models.Client
	query := `select id, client_name, client_surname, birthday,
	                 gender, registration_date, address_id
			  from clients
			  where client_name = $1, client_surname = $2`
	err := r.db.GetContext(ctx, &model, query, fullname.Name, fullname.Surname)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *ClientRepo) Create(ctx context.Context, model *models.Client) error {
	query := `insert into clients (client_name, client_surname, birthday,
			                       gender, registration_date, address_id)
			  values ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query,
		model.ClientName, model.ClientSurname, model.Birthday,
		model.Gender, model.RegistrationDate, model.AddressId)
	return err
}

func (r *ClientRepo) Update(ctx context.Context, model *models.Client) error {
	query := `update clients set address_id = $1 where id = $2`
	_, err := r.db.ExecContext(ctx, query, model.AddressId, model.Id)
	return err
}

func (r *ClientRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from clients where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
