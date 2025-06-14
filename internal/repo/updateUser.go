package repo

import (
	"context"
	"fmt"

	"github.com/MadeSimplest/users/internal/queries/userQuery"

	"github.com/MadeSimplest/core"
)

func UpdateUser(userID string, password string, email string, userStatus string, roleID int64, username string) error {
	// Get database connection
	dbConn := core.DB()
	// fmt.Println(userID, password, email, userStatus, *roleID)

	// Begin transaction
	tx, err := dbConn.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer rollback in case of failure
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Create query with transaction context
	query := userQuery.New(tx)
	ctx := context.Background()

	// Prepare update parameters
	updateUserParams := userQuery.UpdateUserParams{
		ID:       userID,
		Username: username,
		Email:    email,
		Password: password,
	}

	// Update user status (within transaction)
	if _, err := query.UpdateUserStatus(ctx, userQuery.UpdateUserStatusParams{
		MetaValue: userStatus,
		UserID:    userID,
	}); err != nil {
		return fmt.Errorf("failed to update user status: %w", err)
	}

	// Update user role (within transaction)
	if _, err := query.UpdateUserRole(ctx, userQuery.UpdateUserRoleParams{
		UserID: userID,
		RoleID: roleID,
	}); err != nil {
		return fmt.Errorf("failed to update user role: %w", err)
	}

	// Update user (within transaction)
	if _, err := query.UpdateUser(ctx, updateUserParams); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	// Commit transaction if everything succeeded
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func AddUserMeta(userMeta userQuery.AddUserMetaParams) {

	query := userQuery.New(core.DB())

	query.AddUserMeta(context.Background(), userMeta)

}
