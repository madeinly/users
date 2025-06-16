package repo

import (
	"context"
	"fmt"

	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/models"
	"github.com/madeinly/users/internal/queries/userQuery"

	"github.com/madeinly/core"

	"github.com/google/uuid"
)

func CreateUser(username string, email string, password string, roleID models.RoleID, status string) (string, error) {
	ctx := context.Background()

	tx, err := core.DB().BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := userQuery.New(tx)

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		RoleID:   roleID,
		Status:   status,
	}

	user.ID = uuid.New().String()

	hashedPass, err := auth.HashPassword(user.Password)

	if err != nil {
		return "", fmt.Errorf("password hashing failed: %w", err)
	}

	user.Password = hashedPass

	if _, err := query.CreateUser(ctx, userQuery.CreateUserParams{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	if err := query.AddUserRole(ctx, userQuery.AddUserRoleParams{
		UserID: user.ID,
		RoleID: int64(user.RoleID),
	}); err != nil {
		return "", fmt.Errorf("failed to add user role: %w", err)
	}

	if _, err := query.AddUserMeta(ctx, userQuery.AddUserMetaParams{
		UserID:    user.ID,
		MetaKey:   "user_status",
		MetaValue: user.Status,
	}); err != nil {
		return "", fmt.Errorf("failed to add user meta: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return user.ID, nil
}
