package http

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/google/uuid"
// 	"github.com/madeinly/users/internal/auth"
// 	"github.com/madeinly/users/internal/repository"
// 	"github.com/madeinly/users/internal/user"
// )

// func Authenticate(w http.ResponseWriter, r *http.Request) {
// 	uv := user.NewUserValidator()

// 	if err := r.ParseForm(); err != nil {
// 		uv.AddError("BadRequest", "Could not parse the form", user.PropUserForm)
// 		uv.RespondErrors(w)
// 		return
// 	}

// 	email := r.FormValue(user.PropUserEmail)
// 	password := r.FormValue(user.PropUserPassword)

// 	uv.ValidEmail(email)
// 	uv.ValidPassword(password)

// 	if uv.HasErrors() {
// 		uv.RespondErrors(w)
// 		return
// 	}

// 	repo := repository.NewUserRepo()
// 	isValid, userID := repo.ValidateCredentials(email, password)

// 	if !isValid {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	fmt.Println("userID", userID)
// 	user := repo.GetByID(userID)
// 	fmt.Println("user", user)
// 	roleID := strconv.FormatInt(int64(user.RoleID), 10)

// 	// Always generate new token (better security)
// 	expiresAt := time.Now().Add(2 * time.Hour)
// 	authToken, err := auth.GenerateToken(userID, uuid.NewString(), roleID, expiresAt)
// 	if err != nil {
// 		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
// 		return
// 	}

// 	// Update or create session
// 	session := repo.GetSessionByUserID(userID)
// 	if session.IsEmpty() {
// 		session = repo.CreateUserSession(userID)
// 	} else {
// 		session = repo.UpdateUserSession(userID, authToken, expiresAt)
// 	}

// 	// Return response
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(auth.ResToken{
// 		Token:     authToken,
// 		ExpiresAt: expiresAt.Format("2006-01-02T15:04:05-07:00"),
// 	}); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func ValidateToken(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"error": "Only POST method is allowed",
// 		})
// 		return
// 	}

// 	authHeader := r.Header.Get("Authorization")
// 	if authHeader == "" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"error": "Authorization header is missing",
// 		})
// 		return
// 	}

// 	if !strings.HasPrefix(authHeader, "Bearer ") {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"error": "Invalid authorization format. Expected 'Bearer <token>'",
// 		})
// 		return
// 	}

// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 	claims, err := auth.ParseToken(tokenString)

// 	if err != nil {
// 		var statusCode int
// 		var errorMsg string

// 		switch err {
// 		case jwt.ErrSignatureInvalid:
// 			statusCode = http.StatusUnauthorized
// 			errorMsg = "Invalid token signature"
// 		case jwt.ErrTokenExpired:
// 			statusCode = http.StatusUnauthorized
// 			errorMsg = "Token has expired"
// 		case jwt.ErrTokenNotValidYet:
// 			statusCode = http.StatusUnauthorized
// 			errorMsg = "Token is not valid yet"
// 		default:
// 			statusCode = http.StatusBadRequest
// 			errorMsg = "Invalid token"
// 		}

// 		w.WriteHeader(statusCode)
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"error": errorMsg,
// 		})
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message": "Token is valid",
// 		"claims":  claims,
// 	})
// }
