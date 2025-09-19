package flows

import (
	"context"

	"github.com/madeinly/users/internal/features/user"
)

type UpdateUserParams struct {
	ID       string
	Username string
	Email    string
	Status   string
	Password string
	Role     string
}

func UpdateUser(ctx context.Context, params UpdateUserParams) error {

	repoParams := user.UpdateUserParams{
		ID:       params.ID,
		Username: params.Username,
		Email:    params.Email,
		Status:   params.Status,
		Password: params.Password,
		Role:     params.Role,
	}

	err := user.Update(ctx, repoParams)

	if err != nil {
		return err
	}

	return nil
}
