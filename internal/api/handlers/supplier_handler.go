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

// GetAllSuppliers godoc
// @Summary Get all suppliers
// @Description Get all supplier
// @Tags suppliers
// @Produce json
// @Success 200 {object} dto.SupplierResponse "Suppliers array"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /suppliers [get]
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

// GetSupplier godoc
// @Summary Get supplier by name and surname
// @Description Get supplier details by name and surname
// @Tags suppliers
// @Produce json
// @Param name path string true "name" format(string)
// @Param surname path string true "surname" format(string)
// @Success 200 {object} dto.SupplierResponse "Supplier found"
// @Failure 400 {object} utils.ErrorResponse "Name or surname are not specified"
// @Failure 404 {object} utils.ErrorResponse "Supplier not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /suppliers/{name}/{surname} [get]
func (h SupplierHandler) GetSupplierById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	supplier, err := h.supplierService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Supplier not found")
		return
	}

	var response *dto.SupplierResponse
	response = dto.ModelToSupplierResponse(supplier)
	utils.SendJSON(w, http.StatusOK, response)
}

// CreateSupplier godoc
// @Summary Create a new supplier
// @Description Create a new supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param request body dto.CreateSupplierRequest true "Supplier data"
// @Success 200 {object} dto.SupplierResponse "Supplier created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or validation error"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /suppliers [post]
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

// UpdateSupplierAddress godoc
// @Summary Update supplier
// @Description Update supplier address
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID" format(uuid)
// @Param request body dto.UpdateSupplierAddressRequest true "Supplier address to update"
// @Success 200 {object} dto.SupplierResponse "Supplier address updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or supplier ID"
// @Failure 404 {object} utils.ErrorResponse "Supplier not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /suppliers/{id} [put]
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
		utils.SendError(w, http.StatusNotFound, "Supplier not found")
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

// DeleteSupplier godoc
// @Summary Delete supplier
// @Description Delete supplier by ID
// @Tags suppliers
// @Param id path string true "User ID" format(uuid)
// @Success 200 "Supplier deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid supplier ID"
// @Failure 404 {object} utils.ErrorResponse "Supplier not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /suppliers/{id} [delete]
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
