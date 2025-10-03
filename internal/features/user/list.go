package user

import (
	"context"
	"fmt"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/drivers/sqlite/sqlc"
)

type UserListParams struct {
	Username string
	Status   string
	Role     string
	Offset   int64
	Page     int
	Limit    int64
}

func List(ctx context.Context, params UserListParams) ([]sqlc.User, int64, error) {

	query := sqlc.New(core.DB())

	us, err := query.GetUsers(ctx, sqlc.GetUsersParams{
		Username: params.Username,
		Status:   params.Status,
		Role:     params.Role,
		Offset:   params.Offset,
		Limit:    params.Limit,
	})

	if err != nil {
		return us, 0, err
	}

	countedUsers, err := query.CountFilteredUsers(ctx, sqlc.CountFilteredUsersParams{
		Username: params.Username,
		Status:   params.Status,
		Role:     params.Role,
	})

	fmt.Println("value of counterUsers", countedUsers)

	if err != nil {
		return us, 0, err
	}

	return us, countedUsers, nil
}
