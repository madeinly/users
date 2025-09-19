package flows

import (
	"context"

	"github.com/madeinly/users/internal/features/user"
)

func UnregisterUser(ctx context.Context, userID string) error {

	err := user.Delete(ctx, userID)

	if err != nil {
		return err
	}

	return nil

}
