package dto

import (
	"storeapi/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateClientRequest struct {
	ClientName       string    `json:"client_name" validate:"required,min=1,max=100"`
	ClientSurname    string    `json:"client_surname" validate:"required,min=1,max=100"`
	Birthday         time.Time `json:"birthday" validate:"required,date"`
	Gender           string    `json:"gender" validate:"required,min=1,max=1"`
	RegistrationDate time.Time `json:"registration_date" validate:"required,date"`
	AddressId        string    `json:"address_id" validate:"omitempty,uuid"`
}

type UpdateClientAddressRequest struct {
	AddressId string `json:"address_id" validate:"required,uuid"`
}

type ClientResponse struct {
	Id               uuid.UUID `json:"id"`
	ClientName       string    `json:"client_name"`
	ClientSurname    string    `json:"client_surname"`
	Birthday         time.Time `json:"birthday"`
	Gender           string    `json:"gender"`
	RegistrationDate time.Time `json:"registration_date"`
	AddressId        string    `json:"address_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func ModelToClientResponse(m *models.Client) *ClientResponse {

}

func ModelToClientResponseList(m []*models.Client) []*ClientResponse {

}
