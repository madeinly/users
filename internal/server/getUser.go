package server

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/models"
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

	repo := models.NewRepo(core.DB())

	u := repo.GetByID(user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if u.ID == "" {
		http.Error(w, "No existe usuario con el id: "+user.ID, http.StatusNotFound)
		return
	}

	resUser := models.NewUser()

	resUser.ID = u.ID
	resUser.Username = u.Username
	resUser.Email = u.Email
	resUser.RoleID = u.RoleID
	resUser.Status = u.Status
	resUser.RoleName = u.RoleID.GetRoleName()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
