package repository

import (
	"context"
	"fmt"

	"github.com/madeinly/users/internal/queries/userQuery"
)

func (repo *sqliteRepo) CheckExist(userID string) bool {

	ctx := context.Background()
	query := userQuery.New(repo.db)

	existUser, err := query.UserExists(ctx, userID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return existUser
}
