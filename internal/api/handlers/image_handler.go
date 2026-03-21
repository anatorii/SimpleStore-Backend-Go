package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ImageHandler struct {
	// userService service.UserService
	validate *validator.Validate
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{
		// userService: userService,
		validate: validator.New(),
	}
}

func (h ImageHandler) GetImageById(w http.ResponseWriter, r *http.Request) { return }

func (h ImageHandler) GetProductImageById(w http.ResponseWriter, r *http.Request) { return }

func (h ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) { return }

func (h ImageHandler) UpdateImage(w http.ResponseWriter, r *http.Request) { return }

func (h ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) { return }
