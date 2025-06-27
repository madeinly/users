package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/madeinly/users/internal/models"
	"github.com/madeinly/users/internal/queries/userQuery"
	"github.com/madeinly/users/internal/repo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser()
	err := r.ParseForm()

	if err != nil {
		user.AddError("form", "The form could not be parse")
		user.RespondErrors(w)
		return
	}

	user.AddUsername(models.ParseUserPOST(r, models.PropUserUsername))
	user.AddEmail(models.ParseUserPOST(r, models.PropUserEmail))
	user.AddPassword(models.ParseUserPOST(r, models.PropUserPassword))
	user.AddStatus(models.ParseUserPOST(r, models.PropUserStatus))
	user.AddRoleID(models.ParseUserPOST(r, models.PropUserRoleID))

	if user.HasErrors() {
		user.RespondErrors(w)
		return
	}

	var u userQuery.User

	u = repo.GetUserByUsername(user.Username)
	if u.Username != "" {
		user.AddError("user_username", "el username ya existe")
	}

	u = repo.GetUserByEmail(user.Email)
	if u.Email != "" {
		user.AddError("user_email", "el correo ya existe")
	}

	if user.HasErrors() {
		user.RespondErrors(w)
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
