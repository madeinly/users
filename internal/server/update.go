package server

import (
	"fmt"
	"net/http"

	"github.com/MadeSimplest/users/internal/auth"
	"github.com/MadeSimplest/users/internal/models"
	"github.com/MadeSimplest/users/internal/parser"
	"github.com/MadeSimplest/users/internal/repo"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	v := parser.NewUserParser()
	err := r.ParseForm()

	if err != nil {
		v.AddError("_form", err.Error())
		v.RespondWithErrors(w)
		return
	}

	user := models.User{}

	if _, exists := r.PostForm[string(parser.FormID)]; exists {
		user.ID = v.FormParse(parser.FormID, r).(string)
	}

	u, _ := repo.GetUserByID(user.ID)

	if u.ID == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if _, exists := r.PostForm[string(parser.FormEmail)]; exists {
		user.Email = v.FormParse(parser.FormEmail, r).(string)
	} else {
		user.Email = u.Email
	}

	if _, exists := r.PostForm[string(parser.FormStatus)]; exists {
		user.Status = v.FormParse(parser.FormStatus, r).(string)
	} else {
		user.Status = u.UserStatus
	}

	if _, exists := r.PostForm[string(parser.FormRoleID)]; exists {
		user.RoleID = v.FormParse(parser.FormRoleID, r).(models.RoleID)
	} else {
		user.RoleID = models.RoleID(u.RoleID)
	}

	if _, exists := r.PostForm[string(parser.FormUsername)]; exists {
		user.Username = v.FormParse(parser.FormUsername, r).(string)
	} else {
		user.Username = u.Username
	}

	if _, exists := r.PostForm[string(parser.FormPassword)]; exists {
		user.Password = v.FormParse(parser.FormPassword, r).(string)
		user.Password, err = auth.HashPassword(user.Password)

		if err != nil {
			v.AddError(parser.FormPassword, err.Error())
		}

	} else {
		user.Password = u.Password
	}

	if v.HasErrors() {
		v.RespondWithErrors(w)
		return
	}

	if err := repo.UpdateUser(user.ID, user.Password, user.Email, user.Status, int64(user.RoleID), user.Username); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to update user", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User updated successfully")
}
