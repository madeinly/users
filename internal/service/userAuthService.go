package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/queries/userQuery"
	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrServerFailure      = errors.New("server issue check logs")
	ErrSessionExpired     = errors.New("session has expired")
)

func ValidateCredentials(ctx context.Context, params ValidateCredentialsParams) (string, error) {

	//================================= validation

	errors := core.Validate()

	if params.Email == "" && params.Username == "" {
		errors.Add("authentication", "--agregar-status-code--", "email and username can't both be empty")
	}

	if params.Password == "" {
		errors.Add("authentication", "--agregar-status-code--", "password can't be empty")
	}

	if errors.HasErrors() {
		return "", errors
	}

	if params.Email != "" {
		errors.Validate(params.Email, user.EmailRules)
	}

	if params.Username != "" {
		errors.Validate(params.Username, user.UsernameRules)
	}

	errors.Validate(params.Password, user.PasswordRules)

	if errors.HasErrors() {
		return "", errors
	}

	//=================================

	//NOTE: maybe could not use a repo and instead add all in here

	repo := repository.NewUserRepo()

	user, err := repo.ValidateCredentials(params.Email, params.Password)

	if err != nil {
		return "", ErrInvalidCredentials
	}

	sessionToken := uuid.New().String()

	expirationTime := time.Now().Add(2 * time.Hour)

	token, err := auth.GenerateToken(sessionToken, user.Role)

	if err != nil {
		return "", ErrServerFailure
	}

	// Update or create session
	session := repo.GetSessionByUserID(user.ID)

	if session.ID == "" {
		err = repo.CreateUserSession(ctx, userQuery.CreateSessionParams{
			ID:          uuid.New().String(),
			UserID:      user.ID,
			Token:       sessionToken,
			SessionData: "[]",
			ExpiresAt:   expirationTime.Format("2006-01-02 15:04:05"),
		})

	} else {
		err = repo.UpdateUserSession(user.ID, sessionToken, expirationTime.Format("2006-01-02 15:04:05"))
	}

	if err != nil {
		return "", err
	}

	return token, nil

}

func AuthenticateWithToken(jwtToken string) (bool, error) {

	claims, err := auth.ParseToken(jwtToken)

	if err != nil {
		return false, err
	}

	sessionToken := claims.SessionToken

	q := userQuery.New(core.DB())

	session, err := q.GetSessionByToken(context.Background(), sessionToken)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		core.Log(err.Error(), "could not fetch session from user")
		return false, ErrServerFailure
	}

	exp, err := time.ParseInLocation("2006-01-02 15:04:05", session.ExpiresAt, time.Local)

	if err != nil {
		core.Log(err.Error(), "could not parse the expiration of the session")
		return false, ErrServerFailure
	}

	now := time.Now()
	if now.After(exp) {
		return false, ErrSessionExpired
	}

	return true, nil

}
