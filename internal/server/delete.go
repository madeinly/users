package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/madeinly/users/internal/parser"
	"github.com/madeinly/users/internal/repo"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

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

	userExist := repo.CheckUserExist(userID)

	if !userExist {
		http.Error(w, fmt.Sprintf("the user with id %s does not exist", userID), http.StatusBadRequest)
		return
	}

	err = repo.DeleteUser(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	// Parse JSON
	var request struct {
		UserIDs []string `json:"user_ids"`
	}

	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	v := parser.NewUserParser()

	for _, userID := range request.UserIDs {
		userID := v.ValidateID(userID)

		if v.HasErrors() {
			v.RespondWithErrors(w)
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
