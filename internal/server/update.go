package server

import (
	"fmt"
	"net/http"

	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	uv := user.NewUserValidator()
	err := r.ParseForm()

	if err != nil {
		uv.AddError("BadRequest", err.Error(), user.PropUserForm)
		uv.RespondErrors(w)
		return
	}

	userID := r.FormValue(user.PropUserID)
	user.IdValidation(userID)

	repo := repository.NewUserRepo()

	u := repo.GetByID(userID)

	if u.IsEmpty() {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userEmail := r.FormValue(user.PropUserEmail)
	uv.ValidEmail(userEmail)

	userStatus := r.FormValue(user.PropUserStatus)
	uv.ValidStatus(userStatus)

	userRoleID := r.FormValue(user.PropUserRoleID)
	uv.ValidRoleID(userRoleID)

	userUsername := r.FormValue(user.PropUserUsername)
	uv.ValidUsername(userUsername)

	userPassword := r.FormValue(user.PropUserPassword)
	uv.ValidPassword(userPassword)

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	hashedPassword, err := auth.HashPassword(userPassword)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	if err := repo.Update(repository.UserArgs{
		ID:       userID,
		Username: userUsername,
		Email:    userEmail,
		Status:   userStatus,
		Password: hashedPassword,
		RoleID:   userRoleID,
	}); err != nil {

		fmt.Println(err.Error())
		http.Error(w, "Failed to update user", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User updated successfully")
}
