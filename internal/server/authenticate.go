package server

import (
	"fmt"
	"net/http"

	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {

	uv := user.NewUserValidator()

	if err := r.ParseForm(); err != nil {
		uv.AddError("BadRequest", "Could not parse the form, possible malformation (?)", user.PropUserForm)
		uv.RespondErrors(w)
		return
	}

	email := r.FormValue(user.PropUserEmail)
	password := r.FormValue(user.PropUserPassword)

	uv.ValidEmail(email)
	uv.ValidPassword(password)

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	repo := repository.NewUserRepo()

	isValid, _ := repo.ValidateCredentials(email, password) // _ is userID

	if !isValid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// sessionToken := uuid.New().String()

	/*
		1. Crear sesion con token opaco
		2. generar un jwt que contenga la informacion
		3. enviar el token
	*/

	fmt.Fprint(w, isValid)
}
