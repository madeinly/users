package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MadeSimplest/users/internal/queries/userQuery"

	"github.com/MadeSimplest/core"
)

func GetUsers(username string, roleID int64, status string, limit int64, offset int64) ([]userQuery.GetUsersRow, error) {
	ctx := context.Background()
	query := userQuery.New(core.DB())

	params := userQuery.GetUsersParams{
		Username: "%" + username + "%",
		RoleID: sql.NullInt64{
			Int64: roleID,
			Valid: true,
		},
		Status: status,
		Limit:  limit,
		Offset: offset,
	}

	users, err := query.GetUsers(ctx, params)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return users, nil
}
