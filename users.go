package users

import (
	"context"
	_ "embed"

	"github.com/google/uuid"
	core "github.com/madeinly/core/v1"

	"github.com/madeinly/users/internal/features/auth"
	"github.com/madeinly/users/internal/features/user"
	"github.com/madeinly/users/internal/gateways/cmd"
	"github.com/madeinly/users/internal/gateways/http"
)

var UserMigration = core.Migration{
	Schema: initialSchema,
	Name:   "users",
}

var Feature = core.FeaturePackage{
	Name:      "users",
	Migration: UserMigration,
	Args:      args,
	Setup:     setupUsers,
	Routes:    http.Routes,
	Cmd:       cmd.Execute,
}

//go:embed internal/drivers/sqlite/sqlc_src/initial_schema.sql
var initialSchema string

var args = []core.Arg{
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

	ctx := context.Background()

	_, err := user.Create(ctx, user.CreateUserParams{
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
