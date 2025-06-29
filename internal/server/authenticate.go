package server

import (
	"net/http"

	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/models"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {

	user := models.NewUser()

	if err := r.ParseForm(); err != nil {
		user.AddError(models.PropUserForm, "Could not parse the form, possible malformation (?)")
		user.RespondErrors(w)
		return
	}

	user.AddEmail(models.ParseUserPOST(r, models.PropUserEmail))
	user.AddPassword(models.ParseUserPOST(r, models.PropUserPassword))

	if user.HasErrors() {
		user.RespondErrors(w)
		return
	}

	isValid, _ := auth.ValidateCredentials(user.Email, user.Password) // _ is userID

	if !isValid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	/*
		1. Crear sesion con token opaco
		2. generar un jwt que contenga la informacion
		3. enviar el token
	*/
}

// func ValidateCookie(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("capre_token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			http.Error(w, "Authorization token missing", http.StatusUnauthorized)
// 			return
// 		}
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	claims, err := auth.ValidateToken(cookie.Value)
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			http.Error(w, "Invalid token signature", http.StatusUnauthorized)
// 			return
// 		}
// 		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
// 		return
// 	}

// 	repoUser, err := repo.GetUserByID(claims.UserID)
// 	if err != nil {
// 		http.Error(w, "internal error", http.StatusInternalServerError)
// 	}

// 	if repoUser.ID == "" {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	user := models.User{
// 		ID:       repoUser.ID,
// 		Username: repoUser.Username,
// 		Email:    repoUser.Email,
// 		Status:   repoUser.UserStatus,
// 		RoleID:   models.RoleID(repoUser.RoleID),
// 		RoleName: models.RoleID(repoUser.RoleID).GetRoleName(),
// 	}

// 	w.WriteHeader(http.StatusOK)

// 	// 4. Return the user as JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }
