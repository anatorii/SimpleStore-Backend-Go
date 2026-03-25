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

type SupplierHandler struct {
	supplierService service.SupplierService
	validate        *validator.Validate
}

func NewSupplierHandler(supplierService service.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		supplierService: supplierService,
		validate:        validator.New(),
	}
}

func (h SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	list, err := h.supplierService.GetAll(r.Context())
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response []*dto.SupplierResponse
	response = dto.ModelToSupplierResponseList(list)
	utils.SendJSON(w, http.StatusOK, response)
}

func (h SupplierHandler) GetSupplierById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	supplier, err := h.supplierService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Supplier not found")
		return
	}

	var response *dto.SupplierResponse
	response = dto.ModelToSupplierResponse(supplier)
	utils.SendJSON(w, http.StatusOK, response)
}

func (h SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateSupplierRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	supplier := models.Supplier{
		Name:        request.Name,
		AddressId:   request.AddressId,
		PhoneNumber: request.PhoneNumber,
	}
	err = h.supplierService.Create(r.Context(), &supplier)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}

func (h SupplierHandler) UpdateSupplierAddress(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	var request dto.UpdateSupplierAddressRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	supplier, err := h.supplierService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Supplier not found")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	supplier.AddressId, _ = uuid.Parse(request.AddressId)
	err = h.supplierService.Update(r.Context(), supplier)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}

func (h SupplierHandler) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	err = h.supplierService.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			utils.SendError(w, http.StatusNotFound, "Supplier not found")
			return
		}
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}
