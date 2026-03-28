package repository

import (
	"context"
	"database/sql"
	"errors"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) ClientRepo {
	return &clientRepo{db: db}
}

func (r *clientRepo) GetAll(ctx context.Context) ([]*models.Client, error) {
	var clients []*models.Client
	query := `select id, client_name, client_surname, birthday, 
			         gender, registration_date, address_id
			  from clients`
	err := r.db.Select(&clients, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return clients, nil
}

func (r *clientRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	var model models.Client
	query := `select id, client_name, client_surname, birthday,
	                 gender, registration_date, address_id
			  from clients
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

func (r *clientRepo) GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error) {
	var model models.Client
	query := `select id, client_name, client_surname, birthday,
	                 gender, registration_date, address_id
			  from clients
			  where client_name = $1 and client_surname = $2`
	err := r.db.Get(&model, query, fullname.Name, fullname.Surname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func (r *clientRepo) Create(ctx context.Context, model *models.Client) error {
	query := `insert into clients (client_name, client_surname, birthday,
			                       gender, registration_date, address_id)
			  values ($1, $2, $3, $4, $5, $6)`

	AddressId := uuid.NullUUID{Valid: false}
	if model.AddressId != uuid.Nil {
		AddressId = uuid.NullUUID{UUID: model.AddressId, Valid: true}
	}
	_, err := r.db.Exec(query,
		model.ClientName, model.ClientSurname, model.Birthday,
		model.Gender, model.RegistrationDate, AddressId)
	return err
}

func (r *clientRepo) Update(ctx context.Context, model *models.Client) error {
	query := `update clients set address_id = $1 where id = $2`
	result, err := r.db.Exec(query, model.AddressId, model.Id)
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

func (r *clientRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from clients where id = $1`
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
