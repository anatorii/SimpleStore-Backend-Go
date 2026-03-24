package handlers

import (
	"net/http"
	"storeapi/internal/api/dto"
	"storeapi/internal/domain/models"
	"storeapi/internal/service"
	"storeapi/pkg/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ClientHandler struct {
	clientService service.ClientService
	validate      *validator.Validate
}

func NewClientHandler(clientService service.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
		validate:      validator.New(),
	}
}

func (h ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	return
}

func (h ClientHandler) GetClientByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	if name == "" || surname == "" {
		utils.SendError(w, http.StatusBadRequest, "Name or surname are not specified")
	}

	client, err := h.clientService.GetByName(r.Context(), models.FullName{Name: name, Surname: surname})
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Client not found")
	}

	var response *dto.ClientResponse
	response = dto.ModelToClientResponse(client)
	utils.SendJSON(w, http.StatusOK, response)
}

func (h ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	return
}

func (h ClientHandler) UpdateClientAddress(w http.ResponseWriter, r *http.Request) {
	return
}

func (h ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	return
}
