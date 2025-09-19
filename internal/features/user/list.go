package user

import (
	"context"

	core "github.com/madeinly/core/v1"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

type UsersPage struct {
	Limit int64 `json:"user_limit"`
	Page  int64 `json:"user_page"`
	Total int   `json:"user_total"`
	Users []User
}

type UserListParams struct {
	Username string
	Status   string
	Role     string
	Offset   int64
	Page     int
	Limit    int64
}

func List(ctx context.Context, params UserListParams) ([]sqlc.User, error) {

	query := sqlc.New(core.DB())

	us, err := query.GetUsers(ctx, sqlc.GetUsersParams{
		Username: params.Username,
		Status:   params.Status,
		Role:     params.Role,
		Offset:   params.Offset,
		Limit:    params.Limit,
	})

	if err != nil {
		return us, err
	}

	return us, nil
}
