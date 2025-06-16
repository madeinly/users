package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/madeinly/users/internal/queries/userQuery"

	"github.com/madeinly/core"
)

func GetUserByID(userID string) (userQuery.GetUserRow, error) {
	ctx := context.Background()
	query := userQuery.New(core.DB())

	user, err := query.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userQuery.GetUserRow{}, nil // Return empty user, no error
		}
		return userQuery.GetUserRow{}, err // Real error (e.g., DB failure)
	}

	return user, nil
}

func GetUserByUsername(username string) userQuery.User {
	ctx := context.Background()
	query := userQuery.New(core.DB())

	user, err := query.GetUserByUsername(ctx, username)

	if err != nil {
		return userQuery.User{}
	}

	return user
}

func GetUserByEmail(email string) userQuery.User {
	ctx := context.Background()
	query := userQuery.New(core.DB())

	user, err := query.GetUserByEmail(ctx, email)

	if err != nil {
		return userQuery.User{}
	}

	return user
}

func CheckUserExist(userID string) bool {

	ctx := context.Background()
	query := userQuery.New(core.DB())

	existUser, err := query.UserExists(ctx, userID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return existUser
}
