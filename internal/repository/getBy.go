package repository

import (
	"context"

	"github.com/madeinly/users/internal/queries/userQuery"
	"github.com/madeinly/users/internal/user"
)

func (repo *sqliteRepo) GetByUsername(username string) user.User {
	ctx := context.Background()
	query := userQuery.New(repo.db)

	u, err := query.GetUserByUsername(ctx, username)

	if err != nil {
		return user.User{}
	}

	return user.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		//Status: , no esta trayendo el status [TODO]
		Password: u.Password,
		//Agregar otros props
	}
}

func (repo *sqliteRepo) GetByEmail(email string) user.User {
	ctx := context.Background()
	query := userQuery.New(repo.db)

	u, err := query.GetUserByEmail(ctx, email)

	if err != nil {
		return user.User{}
	}

	return user.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		//Agregar otros props
	}
}
