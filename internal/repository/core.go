package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/queries/userQuery"
	"github.com/madeinly/users/internal/user"
)

func (repo *sqliteRepo) Create(args UserArgs) (string, error) {
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

	args.ID = uuid.New().String()

	hashedPass, err := auth.HashPassword(args.Password)

	if err != nil {
		return "", fmt.Errorf("password hashing failed: %w", err)
	}

	args.Password = hashedPass

	if _, err := query.CreateUser(ctx, userQuery.CreateUserParams{
		ID:       args.ID,
		Username: args.Username,
		Email:    args.Email,
		Password: args.Password,
	}); err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	roleID, _ := strconv.ParseInt(args.RoleID, 10, 64)

	if err := query.AddUserRole(ctx, userQuery.AddUserRoleParams{
		UserID: args.ID,
		RoleID: roleID,
	}); err != nil {
		return "", fmt.Errorf("failed to add user role: %w", err)
	}

	if _, err := query.AddUserMeta(ctx, userQuery.AddUserMetaParams{
		UserID:    args.ID,
		MetaKey:   "user_status",
		MetaValue: args.Status,
	}); err != nil {
		return "", fmt.Errorf("failed to add user meta: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return args.ID, nil
}

func (repo *sqliteRepo) Delete(userID string) error {
	ctx := context.Background()
	q := userQuery.New(repo.db)

	err := q.DeleteUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil

}

func (repo *sqliteRepo) Update(args UserArgs) error {

	dbConn := repo.db

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
		ID:       args.ID,
		Username: args.Username,
		Email:    args.Email,
		Password: args.Password,
	}

	// Update user status (within transaction)
	if _, err := query.UpdateUserStatus(ctx, userQuery.UpdateUserStatusParams{
		MetaValue: args.Status,
		UserID:    args.ID,
	}); err != nil {
		return fmt.Errorf("failed to update user status: %w", err)
	}

	roleID, _ := strconv.ParseInt(args.RoleID, 10, 64)

	// Update user role (within transaction)
	if _, err := query.UpdateUserRole(ctx, userQuery.UpdateUserRoleParams{
		UserID: args.ID,
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

func (repo *sqliteRepo) GetByID(userID string) user.User {

	ctx := context.Background()
	query := userQuery.New(repo.db)

	u, err := query.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}
		}
	}

	user := user.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Status:   u.UserStatus,
		Password: u.Password,
		RoleID:   user.RoleID(u.RoleID),
		RoleName: user.RoleID(u.RoleID).GetRoleName(),
	}

	return user

}

func (repo *sqliteRepo) List(args UserListArgs) []user.User {
	ctx := context.Background()
	query := userQuery.New(repo.db)

	params := userQuery.GetUsersParams{
		Username: args.Username,
		RoleID: sql.NullInt64{
			Int64: args.RoleID,
			Valid: true,
		},
		Status: args.Status,
		Limit:  args.Limit,
		Offset: args.Offset,
	}

	us, err := query.GetUsers(ctx, params)

	if err != nil {
		fmt.Println("the error was", err.Error())
		return nil
	}

	users := make([]user.User, 0, len(us))

	for _, u := range us {
		users = append(users, user.User{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Status:   u.StatusName.String,
			Password: u.Password,
			RoleID:   user.RoleID(u.RoleID),
			RoleName: user.RoleID(u.RoleID).GetRoleName(),
		})
	}

	return users
}
