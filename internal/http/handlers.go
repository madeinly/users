package http

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	var req struct {
		UserID   string `json:"user_id"`       //|
		Username string `json:"user_username"` //| this jsonTags are hardcoded
		RoleID   string `json:"user_roleID"`   //|	Create a developer
		Status   string `json:"user_status"`   //| Thus:
		Email    string `json:"user_page"`     //|	dependency, careful!
		Password string `json:"user_limit"`    //| GoodLuck :)
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := h.UserService.UpdateUser(r.Context(), req.UserID, req.RoleID, req.Status, req.Email, req.Password, req.Username)

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
