package dto

type UpdateAddressRequest struct {
	Country string `json:"country" validate:"required"`
	City    string `json:"city" validate:"required"`
	Street  string `json:"street" validate:"required"`
}
