package dto

import (
	"fmt"
	"storeapi/internal/domain/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateClientRequest struct {
	ClientName       string `json:"client_name" validate:"required,min=1,max=100" example:"Ivan"`
	ClientSurname    string `json:"client_surname" validate:"required,min=1,max=100" example:"Ivanov"`
	Birthday         string `json:"birthday" validate:"required,datetime=2006-01-02" example:"2000-01-02"`
	Gender           string `json:"gender" validate:"required,oneof=M F" example:"M"`
	RegistrationDate string `json:"registration_date" validate:"required,datetime=2006-01-02" example:"2016-01-02"`
}

type ClientResponse struct {
	Id               uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	ClientName       string    `json:"client_name" example:"Ivan"`
	ClientSurname    string    `json:"client_surname" example:"Ivanov"`
	Birthday         string    `json:"birthday" example:"2000-01-02"`
	Gender           string    `json:"gender" example:"M"`
	RegistrationDate string    `json:"registration_date" example:"2020-01-02"`
	AddressId        uuid.UUID `json:"address_id" example:"00000000-0000-0000-0000-000000000000"`
	CreatedAt        time.Time `json:"created_at" example:"0001-01-01T00:00:00Z"`
	UpdatedAt        time.Time `json:"updated_at" example:"0001-01-01T00:00:00Z"`
}

func ModelToClientResponse(m *models.Client) *ClientResponse {
	r := ClientResponse{
		Id:               m.Id,
		ClientName:       m.ClientName,
		ClientSurname:    m.ClientSurname,
		Birthday:         m.Birthday.Format("2006-01-02"),
		Gender:           m.Gender,
		RegistrationDate: m.RegistrationDate.Format("2006-01-02"),
		AddressId:        m.AddressId,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
	return &r
}

func ModelToClientResponseList(lm []*models.Client) []*ClientResponse {
	l := make([]*ClientResponse, 0)
	for _, m := range lm {
		r := ClientResponse{
			Id:               m.Id,
			ClientName:       m.ClientName,
			ClientSurname:    m.ClientSurname,
			Birthday:         m.Birthday.Format("2006-01-02"),
			Gender:           m.Gender,
			RegistrationDate: m.RegistrationDate.Format("2006-01-02"),
			AddressId:        m.AddressId,
			CreatedAt:        m.CreatedAt,
			UpdatedAt:        m.UpdatedAt,
		}
		l = append(l, &r)
	}
	return l
}

func (r *CreateClientRequest) Validate(validate *validator.Validate) error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	birthday := r.GetBirthday()
	regDate := r.GetRegistrationDate()

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// Проверка возраста
	age := today.Year() - birthday.Year()
	if age < 18 {
		return fmt.Errorf("Client must be at least 18 years old")
	}
	if age > 120 {
		return fmt.Errorf("Invalid age")
	}

	// Проверка даты регистрации
	if regDate.After(today) {
		return fmt.Errorf("Registration date cannot be in the future")
	}

	if regDate.Before(birthday) {
		return fmt.Errorf("Registration date cannot be before birthday")
	}

	return nil
}

func (r *CreateClientRequest) GetBirthday() time.Time {
	t, _ := time.Parse("2006-01-02", r.Birthday)
	return t
}

func (r *CreateClientRequest) GetRegistrationDate() time.Time {
	t, _ := time.Parse("2006-01-02", r.RegistrationDate)
	return t
}
