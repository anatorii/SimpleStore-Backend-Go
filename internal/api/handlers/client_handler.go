package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"storeapi/internal/api/dto"
	"storeapi/internal/domain/models"
	"storeapi/internal/service"
	"storeapi/pkg/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

// GetAllClients godoc
// @Summary Get all clients
// @Description Get all client
// @Tags clients
// @Produce json
// @Success 200 {object} dto.ClientResponse "Clients array"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /clients [get]
func (h ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	list, err := h.clientService.GetAll(r.Context())
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response []*dto.ClientResponse
	response = dto.ModelToClientResponseList(list)
	utils.SendJSON(w, http.StatusOK, response)
}

// GetClient godoc
// @Summary Get client by name and surname
// @Description Get client details by name and surname
// @Tags clients
// @Produce json
// @Param name path string true "name" format(string)
// @Param surname path string true "surname" format(string)
// @Success 200 {object} dto.ClientResponse "Client found"
// @Failure 400 {object} utils.ErrorResponse "Name or surname are not specified"
// @Failure 404 {object} utils.ErrorResponse "Client not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /clients/{name}/{surname} [get]
func (h ClientHandler) GetClientByName(w http.ResponseWriter, r *http.Request) {
	fname := models.FullName{
		Name:    chi.URLParam(r, "name"),
		Surname: chi.URLParam(r, "surname"),
	}
	if fname.Name == "" || fname.Surname == "" {
		utils.SendError(w, http.StatusBadRequest, "Name or surname are not specified")
		return
	}

	client, err := h.clientService.GetByName(r.Context(), fname)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if client == nil {
		utils.SendError(w, http.StatusNotFound, "Client not found")
		return
	}

	var response *dto.ClientResponse
	response = dto.ModelToClientResponse(client)
	utils.SendJSON(w, http.StatusOK, response)
}

// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client
// @Tags clients
// @Accept json
// @Produce json
// @Param request body dto.CreateClientRequest true "Client data"
// @Success 200 "Client created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or validation error"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /clients [post]
func (h ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateClientRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err.Error()))
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	client := models.Client{
		ClientName:       request.ClientName,
		ClientSurname:    request.ClientSurname,
		Birthday:         request.GetBirthday(),
		Gender:           request.Gender,
		RegistrationDate: request.GetRegistrationDate(),
		AddressId:        request.AddressId,
	}
	err = h.clientService.Create(r.Context(), &client)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "Client created successfully")
}

// UpdateClientAddress godoc
// @Summary Update client address
// @Description Update client address
// @Tags clients
// @Accept json
// @Produce json
// @Param id path string true "Client Id" format(uuid)
// @Param request body dto.UpdateAddressRequest true "Client address to update"
// @Success 200 "Client address updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or client ID"
// @Failure 404 {object} utils.ErrorResponse "Client not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /clients/{id} [patch]
func (h ClientHandler) UpdateClientAddress(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}

	var request dto.UpdateAddressRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	client, err := h.clientService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if client == nil {
		utils.SendError(w, http.StatusNotFound, "Client not found")
		return
	}

	address := models.Address{
		Country: request.Country,
		City:    request.City,
		Street:  request.Street,
	}
	err = h.clientService.UpdateAddress(r.Context(), client, address)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "Client address updated successfully")
}

// DeleteClient godoc
// @Summary Delete client
// @Description Delete client by ID
// @Tags clients
// @Param id path string true "User ID" format(uuid)
// @Success 200 "Client deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid client ID"
// @Failure 404 {object} utils.ErrorResponse "Client not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /clients/{id} [delete]
func (h ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid client Id")
		return
	}

	err = h.clientService.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			utils.SendError(w, http.StatusNotFound, "Client not found")
			return
		}
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "Client deleted successfully")
}
