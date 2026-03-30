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

type ImageHandler struct {
	imageService   service.ImageService
	productService service.ProductService
	validate       *validator.Validate
}

func NewImageHandler(imageService service.ImageService, productService service.ProductService) *ImageHandler {
	return &ImageHandler{
		imageService:   imageService,
		productService: productService,
		validate:       validator.New(),
	}
}

// GetImage godoc
// @Summary Get image by id
// @Description Get image details by id
// @Tags images
// @Produce json
// @Param id path string true "Image Id" format(uuid)
// @Success 200 {object} dto.ImageResponse "Image found"
// @Failure 400 {object} utils.ErrorResponse "Invalid image Id"
// @Failure 404 {object} utils.ErrorResponse "Image not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images/{id} [get]
func (h ImageHandler) GetImageById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid image Id")
		return
	}

	image, err := h.imageService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if image == nil {
		utils.SendError(w, http.StatusNotFound, "Image not found")
		return
	}

	utils.DownloadImage(w, image.Data)
}

// GetProductImageById godoc
// @Summary Get product image by product id
// @Description Get product image details by product id
// @Tags images
// @Produce json
// @Param id path string true "Product Id" format(uuid)
// @Success 200 {object} dto.ImageResponse "Image found"
// @Failure 400 {object} utils.ErrorResponse "Invalid product Id"
// @Failure 404 {object} utils.ErrorResponse "Image not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images/product/{id} [get]
func (h ImageHandler) GetProductImageById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product Id")
		return
	}

	product, err := h.productService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if product == nil {
		utils.SendError(w, http.StatusBadRequest, "Product not found")
		return
	}

	image, err := h.imageService.GetById(r.Context(), product.ImageId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if image == nil {
		utils.SendError(w, http.StatusNotFound, "Image not found")
		return
	}

	var response *dto.ImageResponse
	response = dto.ModelToImageResponse(image)
	utils.SendJSON(w, http.StatusOK, response)
}

// CreateImage godoc
// @Summary Create a new image
// @Description Create a new image
// @Tags images
// @Accept json
// @Produce json
// @Param request body dto.CreateImageRequest true "Image to create"
// @Success 200 "Image created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or validation error"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images [post]
func (h ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateImageRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(h.validate); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.productService.GetById(r.Context(), request.ProductId)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Product not found")
		return
	}

	image := &models.Image{
		Data: []byte(request.Data),
	}
	err = h.imageService.Create(r.Context(), image, product)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "Image created successfully")
}

// UpdateImage godoc
// @Summary Update image
// @Description Update image
// @Tags images
// @Accept json
// @Produce json
// @Param id path string true "Image Id" format(uuid)
// @Param request body dto.UpdateImageRequest true "Image to update"
// @Success 200 "Image updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or image Id"
// @Failure 404 {object} utils.ErrorResponse "Image not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images/{id} [patch]
func (h ImageHandler) UpdateImage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid image Id")
		return
	}

	var request dto.UpdateImageRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid json body")
		return
	}

	image, err := h.imageService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if image == nil {
		utils.SendError(w, http.StatusNotFound, "Image not found")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	image.Data = []byte(request.Data)
	err = h.imageService.Update(r.Context(), image)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "Image updated successfully")
}

// DeleteImage godoc
// @Summary Delete image
// @Description Delete image by ID
// @Tags images
// @Param id path string true "User ID" format(uuid)
// @Success 200 "Image deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid image ID"
// @Failure 404 {object} utils.ErrorResponse "Image not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images/{id} [delete]
func (h ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid image Id")
		return
	}

	err = h.imageService.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			utils.SendError(w, http.StatusNotFound, "Image not found")
			return
		}
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "Image deleted successfully")
}
