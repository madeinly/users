package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/flows"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		valid, err := flows.AuthenticateWithToken(token)

		if err != nil && errors.Is(err, flows.ErrSessionExpired) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err != nil {
			core.Log(err.Error(), "there was an issue with authenticate with token service")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
