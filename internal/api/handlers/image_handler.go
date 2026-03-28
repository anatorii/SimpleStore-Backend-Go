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
// @Summary Get image by name and surname
// @Description Get image details by name and surname
// @Tags images
// @Produce json
// @Param name path string true "name" format(string)
// @Param surname path string true "surname" format(string)
// @Success 200 {object} dto.ImageResponse "Image found"
// @Failure 400 {object} utils.ErrorResponse "Name or surname are not specified"
// @Failure 404 {object} utils.ErrorResponse "Image not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images/{name}/{surname} [get]
func (h ImageHandler) GetImageById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	image, err := h.imageService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Image not found")
		return
	}

	var response *dto.ImageResponse
	response = dto.ModelToImageResponse(image)
	utils.SendJSON(w, http.StatusOK, response)
}

// GetImage godoc
// @Summary Get image by name and surname
// @Description Get image details by name and surname
// @Tags images
// @Produce json
// @Param name path string true "name" format(string)
// @Param surname path string true "surname" format(string)
// @Success 200 {object} dto.ImageResponse "User found"
// @Failure 400 {object} utils.ErrorResponse "Name or surname are not specified"
// @Failure 404 {object} utils.ErrorResponse "Image not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images/{name}/{surname} [get]
func (h ImageHandler) GetProductImageById(w http.ResponseWriter, r *http.Request) {
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

	image, err := h.imageService.GetById(r.Context(), product.Id)
	if err != nil {
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
// @Param request body dto.CreateImageRequest true "Image data"
// @Success 200 {object} dto.ImageResponse "Image created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or validation error"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images [post]
func (h ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateImageRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	image := models.Image{
		Data:        []byte(request.Data),
		Description: request.Description,
	}
	err = h.imageService.Create(r.Context(), &image)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
}

// UpdateImage godoc
// @Summary Update image
// @Description Update image
// @Tags images
// @Accept json
// @Produce json
// @Param id path string true "Image ID" format(uuid)
// @Param request body dto.UpdateImageRequest true "Image to update"
// @Success 200 {object} dto.ImageResponse "Image updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or image ID"
// @Failure 404 {object} utils.ErrorResponse "Image not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /images/{id} [put]
func (h ImageHandler) UpdateImage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	var request dto.UpdateImageRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	image, err := h.imageService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Image not found")
		return
	}

	if err := h.validate.Struct(request); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	image.Data = []byte(request.Data)
	image.Description = request.Description
	err = h.imageService.Update(r.Context(), image)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, "ok")
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
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
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

	utils.SendJSON(w, http.StatusOK, "ok")
}
