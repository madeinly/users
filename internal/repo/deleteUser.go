package repo

import (
	"context"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/queries/userQuery"
)

func DeleteUser(userID string) error {
	ctx := context.Background()
	q := userQuery.New(core.DB())

	err := q.DeleteUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil

}
