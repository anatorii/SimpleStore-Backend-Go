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

type ProductHandler struct {
	productService service.ProductService
	validate       *validator.Validate
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		validate:       validator.New(),
	}
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all product
// @Tags products
// @Produce json
// @Success 200 {object} dto.ProductResponse "Products array"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products [get]
func (h ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	list, err := h.productService.GetAll(r.Context())
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response []*dto.ProductResponse
	response = dto.ModelToProductResponseList(list)
	utils.SendJSON(w, http.StatusOK, response)

}

// GetProduct godoc
// @Summary Get product by name and surname
// @Description Get product details by name and surname
// @Tags products
// @Produce json
// @Param name path string true "name" format(string)
// @Param surname path string true "surname" format(string)
// @Success 200 {object} dto.ProductResponse "Product found"
// @Failure 400 {object} utils.ErrorResponse "Name or surname are not specified"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products/{name}/{surname} [get]
func (h ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	product, err := h.productService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	var response *dto.ProductResponse
	response = dto.ModelToProductResponse(product)
	utils.SendJSON(w, http.StatusOK, response)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Product data"
// @Success 200 {object} dto.ProductResponse "Product created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or validation error"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products [post]
func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	product := models.Product{
		Name:           request.Name,
		Category:       request.Category,
		Price:          request.Price,
		AvailableStock: request.AvailableStock,
		LastUpdateDate: request.LastUpdateDate,
		SupplierId:     request.SupplierId,
		ImageId:        request.ImageId,
	}
	err = h.productService.Create(r.Context(), &product)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}

// UpdateProductAvailable godoc
// @Summary Update product
// @Description Update product address
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Param request body dto.UpdateProductAvailableRequest true "Product available in stock to update"
// @Success 200 {object} dto.ProductResponse "Product updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or product ID"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products/{id} [put]
func (h ProductHandler) UpdateProductAvailable(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	var request dto.UpdateProductAvailableRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product, err := h.productService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	product.AvailableStock = request.AvailableStock
	err = h.productService.Update(r.Context(), product)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Param id path string true "User ID" format(uuid)
// @Success 200 "Product deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid product ID"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products/{id} [delete]
func (h ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	err = h.productService.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			utils.SendError(w, http.StatusNotFound, "Product not found")
			return
		}
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}
