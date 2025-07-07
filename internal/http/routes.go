package http

import coreModels "github.com/madeinly/core/models"

var handler = NewHandler()

var Routes = []coreModels.Route{
	{
		Type:    "POST",
		Pattern: "/user",
		Handler: handler.CreateUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.CreateUser)),
	},
	{
		Type:    "GET",
		Pattern: "/user",
		Handler: handler.GetUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.GetUser)),
	},
	{
		Type:    "GET",
		Pattern: "/users",
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.GetUsers)),
		Handler: handler.GetUsers,
	},
	{
		Type:    "PATCH",
		Pattern: "/user",
		Handler: handler.UpdateUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.UpdateUser)),
	},
	// {
	// 	Type:    "POST",
	// 	Pattern: "/user/auth",
	// 	Handler: http.HandlerFunc(server.Authenticate),
	// },
	{
		Type:    "DELETE",
		Pattern: "/user",
		Handler: handler.DeleteUser,
		// Handler: server.AuthMiddleware(http.HandlerFunc(server.DeleteUser)),
	},
	// {
	// 	Type:    "POST",
	// 	Pattern: "/user/validate",
	// 	Handler: http.HandlerFunc(server.ValidateToken),
	// },
	// {
	// 	Type:    "DELETE",
	// 	Pattern: "/users",
	// 	Handler: http.HandlerFunc(server.BulkDelete),
	// 	// Handler: server.AuthMiddleware(http.HandlerFunc(server.BulkDelete)),
	// },
	// {
	// 	Type:    "GET",
	// 	Pattern: "/users/check-username",
	// 	Handler: http.HandlerFunc(server.CheckUsername),
	// },
}
