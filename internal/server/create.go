package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MadeSimplest/users/internal/models"
	"github.com/MadeSimplest/users/internal/parser"
	"github.com/MadeSimplest/users/internal/repo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	v := parser.NewUserParser()
	err := r.ParseForm()

	if err != nil {
		v.AddError("_form", err.Error())
		v.RespondWithErrors(w)
		return
	}

	user := models.User{
		Username: v.FormParse(parser.FormUsername, r).(string),
		Email:    v.FormParse(parser.FormEmail, r).(string),
		Password: v.FormParse(parser.FormPassword, r).(string),
		Status:   v.FormParse(parser.FormStatus, r).(string),
		RoleID:   v.FormParse(parser.FormRoleID, r).(models.RoleID),
	}

	if v.HasErrors() {
		v.RespondWithErrors(w)
		return
	}

	u := repo.GetUserByUsername(user.Username)

	if u.Username != "" {
		http.Error(w, "el username ya existe", http.StatusConflict)
		return
	}

	u = repo.GetUserByEmail(user.Email)

	if u.Username != "" {
		http.Error(w, "el correo ya existe", http.StatusConflict)
		return
	}

	uuid, err := repo.CreateUser(
		user.Username,
		user.Email,
		user.Password,
		user.RoleID,
		user.Status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"uuid":    uuid,
		"message": fmt.Sprintf("User %s created successfully", user.Username),
	})
}
