package server

import (
	"encoding/json"
	"net/http"

	"github.com/MadeSimplest/users/internal/models"
	"github.com/MadeSimplest/users/internal/repo"

	"github.com/MadeSimplest/users/internal/parser"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	v := parser.NewUserParser()

	err := r.ParseForm()

	if err != nil {
		v.AddError("_form", err.Error())
		v.RespondWithErrors(w)
		return
	}

	userID := v.FormParse(parser.FormID, r).(string)

	if v.HasErrors() {
		v.RespondWithErrors(w)
		return
	}

	user, err := repo.GetUserByID(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID == "" {
		http.Error(w, "No existe usuario con el id: "+userID, http.StatusNotFound)
		return
	}

	roleID := models.RoleID(user.RoleID)

	var resUser models.User = models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RoleID:   roleID,
		Status:   user.UserStatus,
		RoleName: roleID.GetRoleName(),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
