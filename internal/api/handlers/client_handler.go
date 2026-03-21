package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ClientHandler struct {
	// userService service.UserService
	validate *validator.Validate
}

func NewClientHandler() *ClientHandler {
	return &ClientHandler{
		// userService: userService,
		validate: validator.New(),
	}
}

func (h ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) { return }

func (h ClientHandler) GetClientByName(w http.ResponseWriter, r *http.Request) { return }

func (h ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) { return }

func (h ClientHandler) UpdateClientAddress(w http.ResponseWriter, r *http.Request) { return }

func (h ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) { return }
