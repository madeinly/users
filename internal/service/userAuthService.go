package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/queries/userQuery"
	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func (s *UserService) ValidateCredentials(ctx context.Context, userEmail string, userPassword string) (string, string, user.UserErrors) {
	uc := user.NewUserChecker()

	uc.Email(userEmail)
	uc.Password(userPassword)

	if uc.HasErrors() {
		return "", "", *uc
	}

	repo := repository.NewUserRepo()

	user, err := repo.ValidateCredentials(userEmail, userPassword)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		uc.AddError("db_error", "bad attempt on db user deletion", "db")
		return "", "", *uc
	}

	sessionToken := uuid.New().String()

	tokenExpiration := time.Now().Add(2 * time.Hour)

	token, err := auth.GenerateToken(user.ID, sessionToken, user.Role, tokenExpiration)

	if err != nil {
		uc.AddError("token_error", "problem generating token", "token")
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
		uc.AddError("toke_error", "bad token generation", "token")
		return "", "", *uc
	}

	return token, tokenExpiration.Format("2006-01-02 15:04:05"), nil

}
