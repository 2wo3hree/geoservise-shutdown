package handler

import (
	"encoding/json"
	"geoservise/internal/model"
	"geoservise/internal/service"
	"net/http"
)

type AddressHandler struct {
	Service *service.Service
}

func NewAddressHandler(service *service.Service) *AddressHandler {
	return &AddressHandler{
		Service: service,
	}
}

func (h *AddressHandler) Search(w http.ResponseWriter, r *http.Request) {
	var req model.RequestAddressSearch
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Не удалось выполнить декодирование", http.StatusBadRequest)
		return
	}
	addresses, err := h.Service.Search(r.Context(), req.Query)
	if err != nil {
		http.Error(w, "ошибка поиска", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.ResponseAddress{Addresses: addresses})
}

func (h *AddressHandler) Geocode(w http.ResponseWriter, r *http.Request) {
	var req model.RequestGeocode
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Не удалось выполнить декодирование", http.StatusBadRequest)
		return
	}
	addresses, err := h.Service.Geocode(r.Context(), req.Lat, req.Lng)
	if err != nil {
		http.Error(w, "ошибка поиска", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.ResponseAddress{Addresses: addresses})
}
