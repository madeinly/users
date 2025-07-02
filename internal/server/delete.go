package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	uv := user.NewUserValidator()

	err := r.ParseForm()

	if err != nil {
		uv.AddError("BadRequest", err.Error(), user.PropUserForm)
		uv.RespondErrors(w)
		return
	}

	userID := r.FormValue(user.PropUserID)

	uv.ValidID(userID)

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	repo := repository.NewUserRepo()

	userExist := repo.CheckExist(userID)

	if !userExist {
		http.Error(w, fmt.Sprintf("the user with id %s does not exist", userID), http.StatusBadRequest)
		return
	}

	err = repo.Delete(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s  deleted successfully", userID)

}

func BulkDelete(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var request struct {
		UserIDs []string `json:"user_ids"`
	}

	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	uv := user.NewUserValidator()

	for _, userID := range request.UserIDs {
		uv.ValidID(userID)

		if uv.HasErrors() {
			uv.RespondErrors(w)
			return
		}

		repo := repository.NewUserRepo()

		userExist := repo.CheckExist(userID)

		if !userExist {
			http.Error(w, fmt.Sprintf("the user with id %s does not exist", userID), http.StatusBadRequest)
			continue
		}

		err = repo.Delete(userID)

		if err != nil {
			http.Error(w, fmt.Sprintf("the user %s could not be deleted", userID), http.StatusInternalServerError)
			return
		}

	}

	fmt.Println(request)

	w.WriteHeader((http.StatusOK))
	fmt.Fprint(w, "All users were succesfully deleted")

}
