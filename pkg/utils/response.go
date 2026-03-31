package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error" example:"Internal server error"`
}

type SuccessResponse struct {
	Success string `json:"success" example:"Client data updated"`
}

func SendJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal server error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func SendSuccess(w http.ResponseWriter, message string) {
	SendJSON(w, http.StatusOK, SuccessResponse{Success: message})
}

func SendError(w http.ResponseWriter, code int, message string) {
	SendJSON(w, code, ErrorResponse{Error: message})
}

func DownloadImage(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=data.img")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
