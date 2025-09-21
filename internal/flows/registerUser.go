package flows

import (
	"context"

	"github.com/google/uuid"
	"github.com/madeinly/users/internal/features/user"
)

type RegisteUserParams struct {
	Username string
	Email    string
	Password string
	Role     string
	Status   string
}

func RegisterUser(ctx context.Context, params RegisteUserParams) error {

	_, err := user.Create(ctx, user.CreateUserParams{
		UserID:   uuid.NewString(),
		Username: params.Username,
		Email:    params.Email,
		Password: params.Password,
		Role:     params.Role,
		Status:   params.Status,
	})

	if err != nil {
		return err
	}

	return nil
}
