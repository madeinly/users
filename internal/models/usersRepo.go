package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/queries/userQuery"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo interface {
	GetByID()
	CheckExist()
	GetByEmail()
	GetUserByUsername()
	list()
	Delete()
	Update()
	Create()
}

type SqliteRepo struct {
	db *sql.DB
}

func NewRepo() SqliteRepo {
	return SqliteRepo{
		db: core.DB(),
	}

}

func (repo *SqliteRepo) GetByID(userID string) user {

	ctx := context.Background()
	query := userQuery.New(repo.db)

	u, err := query.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user{}
		}
	}

	user := user{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Status:   u.UserStatus,
		Password: u.Password,
		RoleID:   RoleID(u.RoleID),
		RoleName: RoleID(u.RoleID).GetRoleName(),
	}

	return user
}

func (repo *SqliteRepo) List(username string, roleID int64, status string, limit int64, offset int64) []user {
	ctx := context.Background()
	query := userQuery.New(repo.db)

	params := userQuery.GetUsersParams{
		Username: username,
		RoleID: sql.NullInt64{
			Int64: roleID,
			Valid: roleID != 0,
		},
		Status: status,
		Limit:  limit,
		Offset: offset,
	}

	us, err := query.GetUsers(ctx, params)
	if err != nil {
		fmt.Println(err.Error())
		return nil //Pensar si quiero enviar nil o si quiero enviar un arreglo vacio
	}

	users := make([]user, 0, len(us))

	for _, u := range us {
		users = append(users, user{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Status:   u.StatusName.String,
			Password: u.Password,
			RoleID:   RoleID(u.RoleID),
			RoleName: RoleID(u.RoleID).GetRoleName(),
		})
	}

	return users
}

func (repo *SqliteRepo) Delete(userID string) error {
	ctx := context.Background()
	q := userQuery.New(repo.db)

	err := q.DeleteUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil

}

func (repo *SqliteRepo) Update(userID string, password string, email string, userStatus string, roleID int64, username string) error {

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

func (repo *SqliteRepo) Create(username string, email string, password string, roleID RoleID, status string) (string, error) {
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

	user := user{
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

func (repo *SqliteRepo) GetByUsername(username string) user {
	ctx := context.Background()
	query := userQuery.New(repo.db)

	u, err := query.GetUserByUsername(ctx, username)

	if err != nil {
		return user{}
	}

	return user{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		//Status: , no esta trayendo el status [TODO]
		Password: u.Password,
		//Agregar otros props
	}
}

func (repo *SqliteRepo) GetByEmail(email string) user {
	ctx := context.Background()
	query := userQuery.New(repo.db)

	u, err := query.GetUserByEmail(ctx, email)

	if err != nil {
		return user{}
	}

	return user{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		//Agregar otros props
	}
}

func (repo *SqliteRepo) CheckExist(userID string) bool {

	ctx := context.Background()
	query := userQuery.New(repo.db)

	existUser, err := query.UserExists(ctx, userID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return existUser
}

func (repo *SqliteRepo) ValidateCredentials(email string, password string) (bool, string) {
	ctx := context.Background()

	queries := userQuery.New(core.DB())

	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, ""
		}
		log.Printf("Database error: %v", err)
		return false, ""
	}

	// Compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, ""
	}

	return true, user.ID
}
