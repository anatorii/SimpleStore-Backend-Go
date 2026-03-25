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

func (h ImageHandler) GetImageById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	image, err := h.imageService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Image not found")
		return
	}

	var response *dto.ImageResponse
	response = dto.ModelToImageResponse(image)
	utils.SendJSON(w, http.StatusOK, response)
}

func (h ImageHandler) GetProductImageById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Id is not specified")
		return
	}

	product, err := h.productService.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Product not found")
		return
	}

	image, err := h.imageService.GetById(r.Context(), product.Id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Image not found")
		return
	}

	var response *dto.ImageResponse
	response = dto.ModelToImageResponse(image)
	utils.SendJSON(w, http.StatusOK, response)
}

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
