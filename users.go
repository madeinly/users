package users

import (
	"context"
	_ "embed"

	"github.com/google/uuid"
	coreModels "github.com/madeinly/core/models"

	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/cmd"
	"github.com/madeinly/users/internal/http"
	"github.com/madeinly/users/internal/repository"
)

var UserMigration = coreModels.Migration{
	Schema: initialSchema,
	Name:   "users",
}

var Feature = coreModels.FeaturePackage{
	Name:      "users",
	Migration: UserMigration,
	Args:      args,
	Setup:     setupUsers,
	Routes:    http.Routes,
	Cmd:       cmd.Execute,
}

//go:embed internal/queries/initial_schema.sql
var initialSchema string

var args = []coreModels.Arg{
	{
		Name:        "username",
		Default:     "admin",
		Required:    false,
		Description: "the initial super admin",
	},
	{
		Name:        "email",
		Required:    true,
		Description: "initial email for super admin",
	},
	{
		Name:        "password",
		Required:    true,
		Description: "password for initial super admin",
	},
}

func setupUsers(params map[string]string) error {

	repo := repository.NewUserRepo()
	ctx := context.Background()

	_, err := repo.Create(ctx, repository.CreateUserParams{
		UserID:   uuid.NewString(),
		Username: params["username"],
		Email:    params["email"],
		Password: auth.HashPassword(params["password"]),
		Role:     "role_admin",
		Status:   "active",
	})

	if err != nil {
		return err
	}

	return nil
}
