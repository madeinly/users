package server

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/users/internal/models"
	"github.com/madeinly/users/internal/repo"
)

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	user := models.NewUser()
	err := r.ParseForm()
	if err != nil {
		user.AddError("form", err.Error())
		user.RespondErrors(w)
	}

	user.AddUsername(models.ParseUserGET(r, models.PropUserUsername))

	if user.HasErrors() {
		user.RespondErrors(w)
		return
	}

	u := repo.GetUserByUsername(user.Username)
	if u.Username == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"exists":  false,
			"message": "username is available",
		})
		return
	}

	// Username exists
	json.NewEncoder(w).Encode(map[string]interface{}{
		"exists":  true,
		"message": "username already taken",
	})
}
