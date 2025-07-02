package server

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func CheckUsername(w http.ResponseWriter, r *http.Request) {

	uv := user.NewUserValidator()

	err := r.ParseForm()
	if err != nil {
		uv.AddError("BadRequest", err.Error(), user.PropUserForm)
		uv.RespondErrors(w)
	}

	username := r.URL.Query().Get(user.PropUserUsername)

	uv.ValidUsername(username)

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	repo := repository.NewUserRepo()

	u := repo.GetByUsername(username)

	if u.IsEmpty() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"exists":  false,
			"message": "username is available",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"exists":  true,
		"message": "username already taken",
	})
}
