package server

// Define a private custom type for context keys
// type contextKey int

// Enum-like constants for context keys
const (
// userContextKey contextKey = iota
// Add other context keys here if needed
)

// getUserFromContext safely extracts the user from context
// func getUserFromContext(ctx context.Context) (*models.User, bool) {
// 	user, ok := ctx.Value(userContextKey).(*models.User)
// 	return user, ok
// }

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("capre_token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				http.Error(w, "Authorization token missing", http.StatusUnauthorized)
// 				return
// 			}
// 			http.Error(w, "Bad request", http.StatusBadRequest)
// 			return
// 		}

// 		claims, err := auth.ValidateToken(cookie.Value)
// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
// 				return
// 			}
// 			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
// 			return
// 		}

// 		repoUser, err := repo.GetUserByID(claims.UserID)
// 		if err != nil {
// 			http.Error(w, "internal error", http.StatusInternalServerError)
// 			return
// 		}

// 		if repoUser.ID == "" {
// 			http.Error(w, "User not found", http.StatusNotFound)
// 			return
// 		}

// 		// Store user in context for downstream handlers
// 		ctx := context.WithValue(r.Context(), userContextKey, repoUser)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// AdminMiddleware checks for admin role (RoleID = 1)
// func AdminMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		user, ok := r.Context().Value("user").(*models.User)
// 		if !ok || models.RoleID(user.RoleID) != 1 {
// 			http.Error(w, "Admin access required", http.StatusForbidden)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
