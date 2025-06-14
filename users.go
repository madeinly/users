package users

import (
	_ "embed"
	"net/http"

	coreModels "github.com/MadeSimplest/core/models"
	"github.com/MadeSimplest/users/internal/models"

	"github.com/MadeSimplest/users/internal/cmd"
	"github.com/MadeSimplest/users/internal/queries/userQuery"
	"github.com/MadeSimplest/users/internal/repo"
	"github.com/MadeSimplest/users/internal/server"
)

type api struct {
	Create func(username string, email string, password string, roleID models.RoleID, status string) (string, error)
	Update func(userID string, password string, email string, userStatus string, roleID int64, username string) error
	Get    func(username string, roleID int64, status string, limit int64, offset int64) ([]userQuery.GetUsersRow, error)
	Auth   func(email string, password string) (bool, string)
}

var Api = api{
	repo.CreateUser,
	repo.UpdateUser,
	repo.GetUsers,
	repo.Auth,
}

var UserMigration = coreModels.Migration{
	Schema: initialSchema,
}

var Feature = coreModels.FeaturePackage{
	Name:      "users",
	Migration: UserMigration,
	Setup:     setupUsers,
	Routes:    Routes,
	Cmd:       cmd.Execute,
}

//go:embed internal/initial_schema.sql
var initialSchema string

var Routes = []coreModels.Route{
	{
		Type:    "POST",
		Pattern: "/user",
		Handler: http.HandlerFunc(server.CreateUser),
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.CreateUser)),
	},
	{
		Type:    "GET",
		Pattern: "/user",
		Handler: http.HandlerFunc(server.GetUser),
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.GetUser)),
	},
	{
		Type:    "GET",
		Pattern: "/users",
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.GetUsers)),
		Handler: http.HandlerFunc(server.GetUsers),
	},
	{
		Type:    "PATCH",
		Pattern: "/user",
		Handler: http.HandlerFunc(server.UpdateUser),
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.UpdateUser)),
	},
	{
		Type:    "POST",
		Pattern: "/user/auth",
		Handler: http.HandlerFunc(server.Auth),
	},
	{
		Type:    "DELETE",
		Pattern: "/user",
		Handler: http.HandlerFunc(server.DeleteUser),
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.DeleteUser)),
	},
	{
		Type:    "POST",
		Pattern: "/user/validate",
		Handler: http.HandlerFunc(server.ValidateCookie),
	},
	{
		Type:    "DELETE",
		Pattern: "/users",
		Handler: http.HandlerFunc(server.BulkDelete),
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.BulkDelete)),
	},
}

func setupUsers() error {

	_, err := Api.Create("admin", "admin@example.com", "qwer1234", models.RoleID(1), "active")

	if err != nil {
		return err
	}

	return nil
}
