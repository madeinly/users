package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/flows"
)

func AuthUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := r.FormValue("user_email")
	password := r.FormValue("user_password")
	username := r.FormValue("user_username")

	flowParam := flows.ValidateCredentialsParams{
		Email:    email,
		Username: username,
		Password: password,
	}

	token, err := flows.ValidateCredentials(r.Context(), flowParam)

	if err != nil {
		if errors, ok := core.IsErrors(err); ok {
			if errors.HasErrors() {
				errors.WriteHTTP(w)
				return
			}
		}
	}

	if err != nil && errors.Is(err, flows.ErrInvalidCredentials) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil && errors.Is(err, flows.ErrServerFailure) {
		core.Log("auth user, error on: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
