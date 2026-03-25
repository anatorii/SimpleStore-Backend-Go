package handlers

import (
	"encoding/json"
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

func (h ClientHandler) GetClientByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	if name == "" || surname == "" {
		utils.SendError(w, http.StatusBadRequest, "Name or surname are not specified")
		return
	}

	client, err := h.clientService.GetByName(r.Context(), models.FullName{Name: name, Surname: surname})
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Client not found")
		return
	}

	var response *dto.ClientResponse
	response = dto.ModelToClientResponse(client)
	utils.SendJSON(w, http.StatusOK, response)
}

func (h ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateClientRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	client := models.Client{
		ClientName:       request.ClientName,
		ClientSurname:    request.ClientSurname,
		Birthday:         request.Birthday,
		Gender:           request.Gender,
		RegistrationDate: request.RegistrationDate,
		AddressId:        request.AddressId,
	}
	err = h.clientService.Create(r.Context(), &client)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}

func (h ClientHandler) UpdateClientAddress(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	var request dto.UpdateClientAddressRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	client, err := h.clientService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Client not found")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	client.AddressId, _ = uuid.Parse(request.AddressId)
	err = h.clientService.Update(r.Context(), client)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}

func (h ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	err = h.clientService.Delete(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}
