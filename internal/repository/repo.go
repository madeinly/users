package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/queries/userQuery"
)

func (repo *sqliteRepo) Create(ctx context.Context, param CreateUserParams) (string, error) {

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := userQuery.New(tx)

	if _, err := query.CreateUser(ctx, userQuery.CreateUserParams{
		ID:       param.UserID,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Role:     param.Role,
		Status:   param.Status,
	}); err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return param.UserID, nil
}

func (repo *sqliteRepo) Delete(ctx context.Context, userID string) error {

	tx, err := core.DB().BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	q := userQuery.New(tx)

	err = q.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	err = q.DeleteSession(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	err = q.DeleteMetas(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete metas: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil

}

// func (repo *sqliteRepo) Update(args UserArgs) error {

// 	dbConn := repo.db

// 	// Begin transaction
// 	tx, err := dbConn.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %w", err)
// 	}

// 	// Defer rollback in case of failure
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// Create query with transaction context
// 	query := userQuery.New(tx)
// 	ctx := context.Background()

// 	// Prepare update parameters
// 	updateUserParams := userQuery.UpdateUserParams{
// 		ID:       args.ID,
// 		Username: args.Username,
// 		Email:    args.Email,
// 		Password: args.Password,
// 	}

// 	// Update user status (within transaction)
// 	if _, err := query.UpdateUserStatus(ctx, userQuery.UpdateUserStatusParams{
// 		MetaValue: args.Status,
// 		UserID:    args.ID,
// 	}); err != nil {
// 		return fmt.Errorf("failed to update user status: %w", err)
// 	}

// 	roleID, _ := strconv.ParseInt(args.RoleID, 10, 64)

// 	// Update user role (within transaction)
// 	if _, err := query.UpdateUserRole(ctx, userQuery.UpdateUserRoleParams{
// 		UserID: args.ID,
// 		RoleID: roleID,
// 	}); err != nil {
// 		return fmt.Errorf("failed to update user role: %w", err)
// 	}

// 	// Update user (within transaction)
// 	if _, err := query.UpdateUser(ctx, updateUserParams); err != nil {
// 		return fmt.Errorf("failed to update user: %w", err)
// 	}

// 	// Commit transaction if everything succeeded
// 	if err := tx.Commit(); err != nil {
// 		return fmt.Errorf("failed to commit transaction: %w", err)
// 	}

// 	return nil
// }

func (repo *sqliteRepo) GetByID(ctx context.Context, userID string) (RepoUser, error) {

	query := userQuery.New(repo.db)

	u, err := query.GetUserByID(ctx, userID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {

		return RepoUser{}, err
	}

	s, err := query.GetSessionByUserID(ctx, userID)

	var lastAccessedAt string

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return RepoUser{}, err
		}
		lastAccessedAt = u.CreatedAt
	} else {
		lastAccessedAt = s.LastAccessedAt
	}

	return RepoUser{
		ID:                u.ID,
		Role:              u.ID,
		Username:          u.Username,
		Email:             u.Email,
		Password:          u.Password,
		Status:            u.Status,
		PasswordUpdatedAt: u.PasswordUpdatedAt,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
		UserStatus:        u.Status,
		LastLogin:         lastAccessedAt,
	}, nil

}

func (repo *sqliteRepo) List(ctx context.Context, params UserListParams) ([]userQuery.User, error) {

	query := userQuery.New(repo.db)

	fmt.Println(params)

	us, err := query.GetUsers(ctx, userQuery.GetUsersParams{
		Username: params.Username,
		Status:   params.Status,
		Role:     params.Role,
		Offset:   params.Offset,
		Limit:    params.Limit,
	})

	if err != nil {
		fmt.Println("the error was", err.Error())
		return us, err
	}

	return us, nil
}
