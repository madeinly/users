package server

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	uv := user.NewUserValidator()

	err := r.ParseForm()

	if err != nil {
		uv.AddError("BadRequest", "could not parse the form", user.PropUserForm)
		uv.RespondErrors(w)
		return
	}

	userID := r.URL.Query().Get(user.PropUserID)
	uv.ValidID(userID)

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	repo := repository.NewUserRepo()

	u := repo.GetByID(userID)

	if u.IsEmpty() {
		http.Error(w, "No existe usuario con el id: "+userID, http.StatusNotFound)
		return
	}

	resUser := user.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Status:   u.Status,
		RoleID:   u.RoleID,
		RoleName: u.RoleID.GetRoleName(),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
