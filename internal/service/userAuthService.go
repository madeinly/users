package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/queries/userQuery"
	"github.com/madeinly/users/internal/repository"
)

func CheckCredentials(userEmail string, username string, userPassword string) (bool, error) {

	return true, nil

}

func ValidateCredentials(ctx context.Context, userEmail string, username string, userPassword string) (string, string, error) {

	repo := repository.NewUserRepo()

	user, err := repo.ValidateCredentials(userEmail, userPassword)

	if err != nil {
		return "", "", err
	}

	sessionToken := uuid.New().String()

	tokenExpiration := time.Now().Add(2 * time.Hour)

	token, err := auth.GenerateToken(user.ID, sessionToken, user.Role)

	if err != nil {
		return "", "", err
	}

	// Update or create session
	session := repo.GetSessionByUserID(user.ID)

	if session.ID == "" {
		err = repo.CreateUserSession(ctx, userQuery.CreateSessionParams{
			ID:          uuid.New().String(),
			UserID:      user.ID,
			Token:       token,
			SessionData: "[]",
			ExpiresAt:   tokenExpiration.Format("2006-01-02 15:04:05"),
		})

	} else {
		err = repo.UpdateUserSession(user.ID, token, tokenExpiration.Format("2006-01-02 15:04:05"))
	}

	if err != nil {
		return "", "", err
	}

	return token, tokenExpiration.Format("2006-01-02 15:04:05"), nil

}
