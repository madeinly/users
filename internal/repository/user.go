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

func (repo *sqliteRepo) Update(ctx context.Context, args UpdateUserParams) error {

	dbConn := repo.db

	tx, err := dbConn.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := userQuery.New(tx)

	if args.Username != "" {
		query.UpdateUserUsername(ctx, userQuery.UpdateUserUsernameParams{
			Username: args.Username,
			ID:       args.ID,
		})
	}

	if args.Email != "" {
		query.UpdateUserEmail(ctx, userQuery.UpdateUserEmailParams{
			Email: args.Email,
			ID:    args.ID,
		})
	}

	if args.Status != "" {
		query.UpdateUserStatus(ctx, userQuery.UpdateUserStatusParams{
			Status: args.Status,
			ID:     args.ID,
		})
	}

	if args.Password != "" {
		query.UpdateUserPassword(ctx, userQuery.UpdateUserPasswordParams{
			Password: args.Password,
			ID:       args.ID,
		})
	}

	if args.Role != "" {
		query.UpdateUserRole(ctx, userQuery.UpdateUserRoleParams{
			Role: args.Role,
			ID:   args.ID,
		})
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (repo *sqliteRepo) GetByID(ctx context.Context, userID string) (RepoUser, error) {

	query := userQuery.New(repo.db)

	u, err := query.GetUserByID(ctx, userID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {

		return RepoUser{}, err
	}

	s, err := query.GetSessionByUserID(ctx, userID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return RepoUser{}, err
	}

	return RepoUser{
		ID:                u.ID,
		Role:              u.Role,
		Username:          u.Username,
		Email:             u.Email,
		Password:          u.Password,
		Status:            u.Status,
		PasswordUpdatedAt: u.PasswordUpdatedAt,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
		UserStatus:        u.Status,
		LastLogin:         s.LastAccessedAt,
	}, nil

}

func (repo *sqliteRepo) List(ctx context.Context, params UserListParams) ([]userQuery.User, error) {

	query := userQuery.New(repo.db)

	us, err := query.GetUsers(ctx, userQuery.GetUsersParams{
		Username: params.Username,
		Status:   params.Status,
		Role:     params.Role,
		Offset:   params.Offset,
		Limit:    params.Limit,
	})

	if err != nil {
		return us, err
	}

	return us, nil
}
