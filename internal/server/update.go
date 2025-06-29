package server

import (
	"fmt"
	"net/http"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	user := models.NewUser()
	err := r.ParseForm()

	if err != nil {
		user.AddError("form", err.Error())
		user.RespondErrors(w)
		return
	}

	if _, exists := r.PostForm[string(models.PropUserID)]; exists {
		user.AddID(models.ParseUserPOST(r, models.PropUserID))
	}

	repo := models.NewRepo(core.DB())

	u := repo.GetByID(user.ID)

	if u.ID == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if _, exists := r.PostForm[string(models.PropUserEmail)]; exists {
		user.AddEmail(models.ParseUserPOST(r, models.PropUserEmail))
	} else {
		user.Email = u.Email
	}

	if _, exists := r.PostForm[string(models.PropUserStatus)]; exists {
		user.AddStatus(models.ParseUserPOST(r, models.PropUserStatus))
	} else {
		user.Status = u.Status
	}

	if _, exists := r.PostForm[string(models.PropUserRoleID)]; exists {
		user.AddRoleID(models.ParseUserPOST(r, models.PropUserRoleID))
	} else {
		user.RoleID = models.RoleID(u.RoleID)
	}

	if _, exists := r.PostForm[string(models.PropUserUsername)]; exists {
		user.AddUsername(models.ParseUserPOST(r, models.PropUserUsername))
	} else {
		user.Username = u.Username
	}

	if _, exists := r.PostForm[string(models.PropUserPassword)]; exists {
		user.AddPassword(models.ParseUserPOST(r, models.PropUserPassword))
		user.Password, err = auth.HashPassword(user.Password)

		if err != nil {
			user.AddError(models.PropUserPassword, "Could not store your password")
		}

	} else {
		user.Password = u.Password
	}

	if user.HasErrors() {
		user.RespondErrors(w)
		return
	}

	if err := repo.Update(user.ID, user.Password, user.Email, user.Status, int64(user.RoleID), user.Username); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to update user", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User updated successfully")
}
