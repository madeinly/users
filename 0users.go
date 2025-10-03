package users

import (
	"context"
	_ "embed"

	"github.com/madeinly/core"

	"github.com/madeinly/users/internal/flows"
	"github.com/madeinly/users/internal/gateways/cmd"
	"github.com/madeinly/users/internal/gateways/http"
)

var UserMigration = core.Migration{
	Schema: initialSchema,
	Name:   "users",
}

var Feature = core.Mod{
	Name:        "users",
	Migration:   UserMigration,
	InstallArgs: installArgs,
	Setup:       setupUsers,
	Routes:      http.Routes,
	Cmd:         cmd.Execute,
}

//go:embed internal/drivers/sqlite/queries/migration.sql
var initialSchema string

var installArgs = []core.InstallArg{
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

	err := flows.RegisterUser(ctx, flows.RegisteUserParams{
		Username: params["username"],
		Email:    params["email"],
		Password: params["password"],
		Role:     "role_admin",
		Status:   "active",
	})

	if err != nil {
		return err
	}

	return nil
}
