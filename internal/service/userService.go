package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/madeinly/users/internal/auth"
	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func RegisterUser(ctx context.Context, params RegisteUserParams) error {

	repo := repository.NewUserRepo()

	// Add here validations related to database structure (max limits chars and so on)
	// Evaluate how to use the same error or evaluate if move to another package for errors

	_, err := repo.Create(ctx, repository.CreateUserParams{
		UserID:   uuid.NewString(),
		Username: params.Username,
		Email:    params.Email,
		Password: auth.HashPassword(params.Password),
		Role:     params.Role,
		Status:   params.Status,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func GetUser(ctx context.Context, userID string) (user.User, error) {

	repo := repository.NewUserRepo()

	repoUser, err := repo.GetByID(ctx, userID)

	if err != nil {
		return user.User{}, err

	}

	return user.User{
		ID:         repoUser.ID,
		Username:   repoUser.Username,
		Role:       repoUser.Role,
		Email:      repoUser.Email,
		Password:   repoUser.Password,
		Status:     repoUser.Status,
		CreatedAT:  repoUser.CreatedAt,
		UpdatedtAt: repoUser.UpdatedAt,
		LastLogin:  repoUser.LastLogin,
	}, nil

}

func UnregisterUser(ctx context.Context, userID string) error {

	repo := repository.NewUserRepo()

	err := repo.Delete(ctx, userID)

	if err != nil {
		return err
	}

	return nil

}

// [TODO] study the relationship between page offset and limit and see if there is a better handling for the
// values that the current implementation
func ListUsers(ctx context.Context, params ListUsersParams) (user.UsersPage, error) {

	var repoParams = repository.UserListParams{
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

	repo := repository.NewUserRepo()

	if repoParams.Page == 1 {
		repoParams.Offset = 0
	} else {
		repoParams.Offset = int64(repoParams.Page) * repoParams.Limit
	}

	us, err := repo.List(ctx, repoParams)

	if err != nil {
		return user.UsersPage{}, err
	}

	var users []user.User

	for _, repoUser := range us {

		users = append(users, user.User{
			ID:         repoUser.ID,
			Username:   repoUser.Username,
			Role:       repoUser.Role,
			Email:      repoUser.Email,
			Password:   repoUser.Password,
			Status:     repoUser.Status,
			CreatedAT:  repoUser.CreatedAt,
			UpdatedtAt: repoUser.UpdatedAt,
		})
	}

	return user.UsersPage{
		Limit: repoParams.Limit,
		Page:  int64(repoParams.Page),
		Total: len(users),
		Users: users,
	}, nil

}

func UpdateUser(ctx context.Context, params UpdateUserParams) error {

	repo := repository.NewUserRepo()

	repoParams := repository.UpdateUserParams{
		ID:       params.UserID,
		Username: params.Username,
		Email:    params.Email,
		Status:   params.Status,
		Password: params.Password,
		Role:     params.Role,
	}

	err := repo.Update(ctx, repoParams)

	if err != nil {
		return err
	}

	return nil
}
