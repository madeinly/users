package http

import (
	coreModels "github.com/madeinly/core/models"
)

var Routes = []coreModels.Route{
	{
		Type:    "POST",
		Pattern: "/user",
		Handler: CreateUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.CreateUser)),
	},
	{
		Type:    "GET",
		Pattern: "/user",
		Handler: GetUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.GetUser)),
	},
	{
		Type:    "GET",
		Pattern: "/users",
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.GetUsers)),
		Handler: GetUsers,
	},
	{
		Type:    "PATCH",
		Pattern: "/user",
		Handler: UpdateUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.UpdateUser)),
	},
	{
		Type:    "POST",
		Pattern: "/user/auth",
		Handler: AuthUser,
	},
	{
		Type:    "DELETE",
		Pattern: "/user",
		Handler: DeleteUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.DeleteUser)),
	},
	{
		Type:    "POST",
		Pattern: "/user/validate",
		Handler: ValidateToken,
	},
	// {
	// 	Type:    "GET",
	// 	Pattern: "/users/check-username",
	// 	Handler: http.HandlerFunc(server.CheckUsername),
	// },
}
