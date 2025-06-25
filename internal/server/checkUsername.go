package server

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/users/internal/parser"
	"github.com/madeinly/users/internal/repo"
)

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	v := parser.NewUserParser()
	err := r.ParseForm()
	if err != nil {
		v.AddError("_form", err.Error())
		v.RespondWithErrors(w)
	}

	username := v.FormParse(parser.FormUsername, r).(string)

	if v.HasErrors() {
		v.RespondWithErrors(w)
		return
	}

	u := repo.GetUserByUsername(username)
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
