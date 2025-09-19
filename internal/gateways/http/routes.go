package http

import (
	"net/http"

	core "github.com/madeinly/core/v1"
)

var Routes = []core.Route{
	{
		Type:    "POST",
		Pattern: "/user",
		Handler: Auth(http.HandlerFunc(CreateUser)),
	},
	{
		Type:    "GET",
		Pattern: "/user",
		Handler: Auth(http.HandlerFunc(GetUser)),
	},
	{
		Type:    "GET",
		Pattern: "/users",
		Handler: Auth(http.HandlerFunc(GetUsers)),
	},
	{
		Type:    "PATCH",
		Pattern: "/user",
		Handler: Auth(http.HandlerFunc(UpdateUser)),
	},
	{
		Type:    "POST",
		Pattern: "/user/auth",
		Handler: http.HandlerFunc(AuthUser),
	},
	{
		Type:    "DELETE",
		Pattern: "/user",
		Handler: Auth(http.HandlerFunc(DeleteUser)),
	},
	{
		Type:    "POST",
		Pattern: "/user/validate",
		Handler: http.HandlerFunc(ValidateToken),
	},
}
