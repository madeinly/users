package flows

import (
	"context"

	"github.com/google/uuid"
	"github.com/madeinly/users/internal/features/auth"
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

	password := auth.HashPassword(params.Password)

	err := user.Create(ctx, user.CreateUserParams{
		UserID:   uuid.NewString(),
		Username: params.Username,
		Email:    params.Email,
		Password: password,
		Role:     params.Role,
		Status:   params.Status,
	})

	if err != nil {
		return ErrServerFailure
	}

	return nil
}
