package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/features/user"
	"github.com/madeinly/users/internal/flows"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	bag := core.Validate()

	if err := r.ParseForm(); err != nil {
		bag.Add("form", "bad_request", "looks like it is malformed could not parse")
		bag.WriteHTTP(w)
		return
	}

	username := r.FormValue("user_username")
	email := r.FormValue("user_email")
	password := r.FormValue("user_password")
	role := r.FormValue("user_role")
	status := r.FormValue("user_status")

	bag.Validate(username, user.UsernameRules)
	bag.Validate(email, user.EmailRules)
	bag.Validate(password, user.PasswordRules)
	bag.Validate(role, user.RoleRules)
	bag.Validate(status, user.StatusRules)

	if bag.HasErrors() {
		bag.WriteHTTP(w)
		return
	}

	err := flows.RegisterUser(r.Context(), flows.RegisteUserParams{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
		Status:   status,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	validator := core.Validate()

	validator.Validate(userID, user.UserIDRules)

	if validator.HasErrors() {
		validator.WriteHTTP(w)
		return
	}

	user, err := flows.GetUser(r.Context(), userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	validator := core.Validate()

	validator.Validate(userID, user.UserIDRules)

	if validator.HasErrors() {

		validator.WriteHTTP(w)
		return
	}

	err := flows.UnregisterUser(r.Context(), userID)

	if err != nil {
		//[!TODO] work on standard errors from user service so I know how to act if something goes wrong there
		return
	}

	w.WriteHeader(http.StatusOK)

}

/*
note that user_limit and user_page are string type, this is cause int would case ambiguity for its
null type, rather use string, and check in the validator that is parseable

[todo] remove the pointer and references its ok to have "" as values if param is note sent
*/
func GetUsers(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	var listUserParams flows.ListUsersParams

	//[!TODO] aqui tendria que chequear es si viene vacio "" para no validar, y en el caso de username que puede ser cualquier cosa seria bueno un metodo
	// que valide que no viene nada malicioso

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

	users, err := flows.ListUsers(r.Context(), listUserParams)

	if err != nil {
		//[!TODO] work on standard errors from user service so I know how to act if something goes wrong there
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	validator := core.Validate()

	if err := r.ParseForm(); err != nil {
		validator.Add("form", "bad_request", "could not parse the form")
		validator.WriteHTTP(w)
		return
	}

	userID := r.FormValue("user_id")
	username := r.FormValue("user_username")
	role := r.FormValue("user_role")
	status := r.FormValue("user_status")
	email := r.FormValue("user_email")
	password := r.FormValue("user_password")

	//to select the user (always validated)
	validator.Validate(userID, user.UserIDRules)

	// validate if present, the ones that are updated
	if username != "" {
		validator.Validate(username, user.UsernameRules)
	}

	if role != "" {
		validator.Validate(role, user.RoleRules)
	}

	if status != "" {
		validator.Validate(status, user.StatusRules)
	}

	if email != "" {
		validator.Validate(email, user.EmailRules)
	}

	if password != "" {
		validator.Validate(password, user.PasswordRules)
	}

	var serviceParams = flows.UpdateUserParams{
		ID:       userID,
		Username: username,
		Role:     role,
		Status:   status,
		Email:    email,
		Password: password,
	}

	ctx := r.Context()
	err := flows.UpdateUser(ctx, serviceParams)

	if err != nil {
		//[!TODO] work on standard errors from user service so I know how to act if something goes wrong there
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")

	fmt.Println(authHeader)
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	valid, err := flows.AuthenticateWithToken(token)

	if err != nil && errors.Is(err, flows.ErrSessionExpired) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		core.Log(err.Error(), "there was an issue with authenticate with token service")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
