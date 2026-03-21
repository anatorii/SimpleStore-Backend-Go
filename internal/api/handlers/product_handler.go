package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	// userService service.UserService
	validate *validator.Validate
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		// userService: userService,
		validate: validator.New(),
	}
}

func (h ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) { return }

func (h ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) { return }

func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) { return }

func (h ProductHandler) UpdateProductAvailable(w http.ResponseWriter, r *http.Request) { return }

func (h ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) { return }
