package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/models"
	"github.com/madeinly/users/internal/queries/userQuery"
	"github.com/madeinly/users/internal/repo"

	"github.com/madeinly/core"

	"golang.org/x/crypto/bcrypt"
)

func Auth(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	isValid, userID := validateCredentials(email, password)

	if !isValid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	cookie, err := auth.SetCookie(userID, email, w, r)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"cookie": cookie,
	})
}

func ValidateCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("capre_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Authorization token missing", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	claims, err := auth.ValidateToken(cookie.Value)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "Invalid token signature", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	repoUser, err := repo.GetUserByID(claims.UserID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	if repoUser.ID == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user := models.User{
		ID:       repoUser.ID,
		Username: repoUser.Username,
		Email:    repoUser.Email,
		Status:   repoUser.UserStatus,
		RoleID:   models.RoleID(repoUser.RoleID),
		RoleName: models.RoleID(repoUser.RoleID).GetRoleName(),
	}

	w.WriteHeader(http.StatusOK)

	// 4. Return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func validateCredentials(email string, password string) (bool, string) {
	ctx := context.Background()

	// Initialize queries with your database connection
	queries := userQuery.New(core.DB()) // Assuming db.Connection is *sql.DB

	// Get user by email using sqlc generated query
	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, ""
		}
		log.Printf("Database error: %v", err)
		return false, ""
	}

	// Compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, ""
	}

	return true, user.ID
}
