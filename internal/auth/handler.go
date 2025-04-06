package auth

import (
	"encoding/json"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary RegisterHandler new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body Credentials true "User credentials"
// @Success 201 {string} string "user registered"
// @Failure 400 {string} string "bad request"
// @Router /register [post]
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Не удалось выполнить декодирование", http.StatusBadRequest)
		return
	}

	if err := RegisterUser(creds.Username, creds.Password); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered"})
}

// @Summary LoginHandler user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body Credentials true "User credentials"
// @Success 200 {object} TokenResponse
// @Failure 400 {string} string "bad request"
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Не удалось выполнить декодирование", http.StatusBadRequest)
		return
	}

	if err := AuthenticateUser(creds.Username, creds.Password); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{"sub": creds.Username})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}
