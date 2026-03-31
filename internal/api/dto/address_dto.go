package dto

type UpdateAddressRequest struct {
	Country string `json:"country" validate:"required,min=1,max=100" example:"Россия"`
	City    string `json:"city" validate:"required,min=1,max=100" example:"Москва"`
	Street  string `json:"street" validate:"required,min=1,max=100" example:"Тверская"`
}
