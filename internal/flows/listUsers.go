package flows

import (
	"context"
	"strconv"

	"github.com/madeinly/users/internal/features/user"
)

type ListUsersParams struct {
	Username *string
	Role     *string
	Status   *string
	Page     *string
	Limit    *string
}

type UsersPage struct {
	Limit int64 `json:"user_limit"`
	Page  int64 `json:"user_page"`
	Total int   `json:"user_total"`
	Users []User
}

type User struct {
	ID                string `json:"user_id"`
	Role              string `json:"user_role"`
	Username          string `json:"user_username"`
	Email             string `json:"user_email"`
	Password          string `json:"-"`
	Status            string `json:"user_status"`
	PasswordUpdatedAt string `json:"-"`
	CreatedAt         string `json:"user_createdAt"`
	UpdatedAt         string `json:"user_updatedAt"`
	LastLogin         string `json:"user_lastLoginAT"`
}

// [TODO] study the relationship between page offset and limit and see if there is a better handling for the
// values that the current implementation
func ListUsers(ctx context.Context, params ListUsersParams) (UsersPage, error) {

	var repoParams = user.UserListParams{
		Limit: 10,
		Page:  1,
	}

	if params.Username != nil {
		repoParams.Username = *params.Username
	}

	if params.Role != nil {
		repoParams.Role = *params.Role
	}

	if params.Status != nil {
		repoParams.Status = *params.Status
	}

	if params.Limit != nil {
		limit, _ := strconv.ParseInt(*params.Limit, 10, 64)
		repoParams.Limit = limit
	}

	if params.Page != nil {
		page, _ := strconv.ParseInt(*params.Page, 10, 64)
		repoParams.Page = int(page)
	}

	if repoParams.Page == 1 {
		repoParams.Offset = 0
	} else {
		repoParams.Offset = int64(repoParams.Page) * repoParams.Limit
	}

	us, err := user.List(ctx, repoParams)

	if err != nil {
		return UsersPage{}, err
	}

	var users []User

	for _, repoUser := range us {

		users = append(users, User{
			ID:        repoUser.ID,
			Username:  repoUser.Username,
			Role:      repoUser.Role,
			Email:     repoUser.Email,
			Password:  repoUser.Password,
			Status:    repoUser.Status,
			CreatedAt: repoUser.CreatedAt,
			UpdatedAt: repoUser.UpdatedAt,
		})
	}

	return UsersPage{
		Limit: repoParams.Limit,
		Page:  int64(repoParams.Page),
		Total: len(users),
		Users: users,
	}, nil

}
