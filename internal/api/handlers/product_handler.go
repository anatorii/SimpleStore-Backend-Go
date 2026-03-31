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
// @Success 200 {object} []dto.ProductResponse "Products array"
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
// @Summary Get product by id
// @Description Get product details by id
// @Tags products
// @Produce json
// @Param id path string true "Product Id" format(uuid) example(00000000-0000-0000-0000-000000000000) default(00000000-0000-0000-0000-000000000000)
// @Success 200 {object} dto.ProductResponse "Product found"
// @Failure 400 {object} utils.ErrorResponse "Invalid product Id"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products/{id} [get]
func (h ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product Id")
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
// @Success 200 {object} utils.SuccessResponse "Product created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or validation error"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products [post]
func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid json body")
		return
	}

	if err := request.Validate(h.validate); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	product := models.Product{
		Name:           request.Name,
		Category:       request.Category,
		Price:          request.Price,
		AvailableStock: request.AvailableStock,
		LastUpdateDate: request.GetLastUpdateDate(),
		SupplierId:     request.GetSupplierId(),
		ImageId:        request.GetImageId(),
	}
	err = h.productService.Create(r.Context(), &product)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(w, "Product created successfully")
}

// UpdateProductAvailable godoc
// @Summary Update product available in stock
// @Description Update product address available in stock
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid) example(00000000-0000-0000-0000-000000000000) default(00000000-0000-0000-0000-000000000000)
// @Param request body dto.UpdateProductAvailableRequest true "Product available in stock to update"
// @Success 200 {object} utils.SuccessResponse "Product available updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or product ID"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products/{id}/available [patch]
func (h ProductHandler) UpdateProductAvailable(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var request dto.UpdateProductAvailableRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid json body")
		return
	}

	product, err := h.productService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product.AvailableStock = request.AvailableStock
	err = h.productService.Update(r.Context(), product)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(w, "Product available updated successfully")
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Param id path string true "User ID" format(uuid) example(00000000-0000-0000-0000-000000000000) default(00000000-0000-0000-0000-000000000000)
// @Success 200 {object} utils.SuccessResponse "Product deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid product Id"
// @Failure 404 {object} utils.ErrorResponse "Product not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /products/{id} [delete]
func (h ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product Id")
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

	utils.SendSuccess(w, "Product deleted successfully")
}
