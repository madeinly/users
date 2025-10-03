package flows

import (
	"context"

	"github.com/madeinly/users/internal/features/user"
)

func GetUser(ctx context.Context, userID string) (User, error) {

	u, err := user.GetByID(ctx, userID)

	if err != nil {
		return User{}, err
	}

	return User{
		ID:                u.ID,
		Role:              u.Role,
		Username:          u.Username,
		Email:             u.Email,
		Password:          u.Password,
		Status:            u.Password,
		PasswordUpdatedAt: u.PasswordUpdatedAt,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
		LastLogin:         "",
	}, err

}
