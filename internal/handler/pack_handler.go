package handler

import (
	"encoding/json"
	"net/http"

	"gymshark/packcalculator/internal/model"
)

type PackHandler struct {
	packService model.PackService
}

func NewPackHandler(packService model.PackService) *PackHandler {
	return &PackHandler{packService: packService}
}

func (h *PackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request model.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.OrderAmount <= 0 {
		http.Error(w, "Order amount must be positive", http.StatusBadRequest)
		return
	}

	response := h.packService.CalculatePacks(request.OrderAmount)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
