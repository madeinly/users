package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/madeinly/users/internal/models"
	"github.com/madeinly/users/internal/repo"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	user := models.NewUser()

	err := r.ParseForm()

	if err != nil {
		user.AddError("_form", err.Error())
		user.RespondErrors(w)
		return
	}

	user.AddID(models.ParseUserPOST(r, models.PropUserID))

	if user.HasErrors() {
		user.RespondErrors(w)
		return
	}

	userExist := repo.CheckUserExist(user.ID)

	if !userExist {
		http.Error(w, fmt.Sprintf("the user with id %s does not exist", user.ID), http.StatusBadRequest)
		return
	}

	err = repo.DeleteUser(user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s  deleted successfully", user.ID)

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

	user := models.NewUser()

	for _, userID := range request.UserIDs {
		user.AddID(userID)

		if user.HasErrors() {
			user.RespondErrors(w)
			return
		}

		userExist := repo.CheckUserExist(userID)

		if !userExist {
			http.Error(w, fmt.Sprintf("the user with id %s does not exist", userID), http.StatusBadRequest)
			continue
		}

		err = repo.DeleteUser(userID)

		if err != nil {
			http.Error(w, fmt.Sprintf("the user %s could not be deleted", userID), http.StatusInternalServerError)
			return
		}

	}

	fmt.Println(request)

	w.WriteHeader((http.StatusOK))
	fmt.Fprint(w, "All users were succesfully deleted")

}
