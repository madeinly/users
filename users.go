package users

import (
	"context"
	_ "embed"

	"github.com/google/uuid"
	"github.com/madeinly/core"
	coreModels "github.com/madeinly/core/models"

	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/cmd"
	"github.com/madeinly/users/internal/http"
	"github.com/madeinly/users/internal/repository"
)

var UserMigration = coreModels.Migration{
	Schema: initialSchema,
}

var Feature = coreModels.FeaturePackage{
	Name:      "users",
	Migration: UserMigration,
	Setup:     setupUsers,
	Routes:    http.Routes,
	Cmd:       cmd.Execute,
}

//go:embed internal/queries/initial_schema.sql
var initialSchema string

func setupUsers() error {

	repo := repository.NewUserRepo()
	ctx := context.Background()

	settings := core.Settings()
	username := settings.User
	password := settings.Password

	_, err := repo.Create(ctx, repository.CreateUserParams{
		UserID:   uuid.NewString(),
		Username: username,
		Email:    "change@email.com",
		Password: auth.HashPassword(password),
		Role:     "role_admin",
		Status:   "active",
	})

	if err != nil {
		return err
	}

	return nil
}
