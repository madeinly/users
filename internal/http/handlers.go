package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/service"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := h.UserService.RegisterUser(r.Context(), service.RegisteUserParams{
		Username: r.FormValue("user_username"),
		Email:    r.FormValue("user_email"),
		Password: r.FormValue("user_password"),
		Role:     r.FormValue("user_role"),
		Status:   r.FormValue("user_status"),
	})

	if err != nil {

		fmt.Println(err)

		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	user, err := h.UserService.GetUser(r.Context(), userID)

	if err != nil {
		respondError(w, err)
		return
	}

	if user.IsEmpty() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	err := h.UserService.UnregisterUser(r.Context(), userID)

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

}

/*
note that user_limit and user_page are string type, this is cause int would case ambiguity for its
null type, rather use string, and check in the validator that is parseable

[todo] remove the pointer and references its ok to have "" as values if param is note sent
*/
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	var listUserParams service.ListUsersParams

	if _, exists := queryParams["user_username"]; exists {
		username := queryParams.Get("user_username")
		listUserParams.Username = &username
	}

	if _, exists := queryParams["user_role"]; exists {
		role := queryParams.Get("user_role")
		listUserParams.Role = &role
	}

	if _, exists := queryParams["user_status"]; exists {
		status := queryParams.Get("user_status")
		listUserParams.Status = &status
	}

	if _, exists := queryParams["user_page"]; exists {
		page := queryParams.Get("user_page")
		listUserParams.Page = &page
	}

	if _, exists := queryParams["user_limit"]; exists {
		limit := queryParams.Get("user_limit")
		listUserParams.Limit = &limit
	}

	users, err := h.UserService.ListUsers(r.Context(), listUserParams)

	if err != nil {
		respondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	var serviceParams = service.UpdateUserParams{
		UserID:   r.FormValue("user_id"),
		Username: r.FormValue("user_username"),
		Role:     r.FormValue("user_role"),
		Status:   r.FormValue("user_status"),
		Email:    r.FormValue("user_email"),
		Password: r.FormValue("user_password"),
	}

	err := h.UserService.UpdateUser(r.Context(), serviceParams)

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, expiration, err := h.UserService.ValidateCredentials(r.Context(), r.FormValue("user_email"), r.FormValue("user_password"))

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token, "expiresAt": expiration})
}

func (h *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	_, err := auth.ParseToken(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
