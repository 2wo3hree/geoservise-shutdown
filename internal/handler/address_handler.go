package handler

import (
	"encoding/json"
	"geoservise-jwt/internal/model"
	"geoservise-jwt/internal/service"
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

// @Summary Search address
// @Description Get city info by address query
// @Tags address
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body model.RequestAddressSearch true "query address"
// @Success 200 {object} model.ResponseAddress
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Failure 403 {string} string "Forbidden"
// @Router /address/search [post]
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

// @Summary Geocode by coordinates
// @Description Get city info by latitude and longitude
// @Tags address
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body model.RequestGeocode true "coordinates"
// @Success 200 {object} model.ResponseAddress
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Failure 403 {string} string "Forbidden"
// @Router /address/geocode [post]
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
