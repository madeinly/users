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

	_, err := repo.Create(ctx, repository.CreateUserParams{
		UserID:   uuid.NewString(),
		Username: "admin",
		Email:    "example@email.com",
		Password: auth.HashPassword("qwer1234"),
		Role:     "role_admin",
		Status:   "active",
	})

	if err != nil {
		return err
	}

	return nil
}
