package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	uv := user.NewUserValidator()
	err := r.ParseForm()

	if err != nil {
		uv.AddError("badRequest", "The form could not be parse", user.PropUserForm)
		uv.RespondErrors(w)
		return
	}

	username := r.FormValue(user.PropUserUsername)
	email := r.FormValue(user.PropUserEmail)
	password := r.FormValue(user.PropUserPassword)
	status := r.FormValue(user.PropUserStatus)
	roleID := r.FormValue(user.PropUserRoleID)

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	repo := repository.NewUserRepo()

	u := repo.GetByUsername(username)

	if !u.IsEmpty() {
		uv.AddError("conflict", "el username ya existe", user.PropUserUsername)
	}

	u = repo.GetByEmail(email)
	if !u.IsEmpty() {
		uv.AddError("conflict", "el correo ya existe", user.PropUserEmail)
	}

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	uuid, err := repo.Create(repository.UserArgs{
		Username: username,
		Email:    email,
		Status:   status,
		RoleID:   roleID,
		Password: password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"uuid":    uuid,
		"message": fmt.Sprintf("User %s created successfully", username),
	})
}
