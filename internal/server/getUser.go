package server

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/users/internal/models"
	"github.com/madeinly/users/internal/repo"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	user := models.NewUser()

	err := r.ParseForm()

	if err != nil {
		user.AddError("form", "could not parse the form")
		user.RespondErrors(w)
		return
	}

	user.AddID(models.ParseUserGET(r, models.PropUserID))

	if user.HasErrors() {
		user.RespondErrors(w)
		return
	}

	u, err := repo.GetUserByID(user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if u.ID == "" {
		http.Error(w, "No existe usuario con el id: "+user.ID, http.StatusNotFound)
		return
	}

	roleID := models.RoleID(user.RoleID)

	var resUser models.User = models.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		RoleID:   roleID,
		Status:   u.UserStatus,
		RoleName: roleID.GetRoleName(),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
