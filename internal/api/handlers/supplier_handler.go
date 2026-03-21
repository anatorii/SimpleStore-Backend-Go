package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type SupplierHandler struct {
	// userService service.UserService
	validate *validator.Validate
}

func NewSupplierHandler() *SupplierHandler {
	return &SupplierHandler{
		// userService: userService,
		validate: validator.New(),
	}
}

func (h SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) { return }

func (h SupplierHandler) GetSupplierById(w http.ResponseWriter, r *http.Request) { return }

func (h SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) { return }

func (h SupplierHandler) UpdateSupplierAddress(w http.ResponseWriter, r *http.Request) { return }

func (h SupplierHandler) DeleteSupplier(w http.ResponseWriter, r *http.Request) { return }
