package flows

import (
	"context"

	"github.com/madeinly/users/internal/features/user"
)

func GetUser(ctx context.Context, userID string) (user.User, error) {

	u, err := user.GetByID(ctx, userID)

	if err != nil {
		return user.User{}, err
	}

	return u, err

}
