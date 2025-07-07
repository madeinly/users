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

type UserService struct{}

func (s *UserService) RegisterUser(ctx context.Context, params RegisteUserParams) user.UserErrors {

	uc := user.NewUserChecker()

	uc.Username(params.Username)
	uc.Email(params.Email)
	uc.Password(params.Password)
	uc.Status(params.Status)

	if uc.HasErrors() { // required to use has error as the object is always initialized as 0 to be able to append errors
		return *uc
	}

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
		uc.AddError("db_error", "bad attempt on db user creation", "db")
		return *uc
	}

	return nil
}

func (s *UserService) GetUser(ctx context.Context, userID string) (user.User, []*user.UserError) {

	uc := user.NewUserChecker()

	uc.UserID(userID)

	if uc.HasErrors() {
		return user.User{}, *uc
	}

	repo := repository.NewUserRepo()

	repoUser, err := repo.GetByID(ctx, userID)

	if err != nil {
		uc.AddError("db_error", "bad attempt on db user get", "db")
		return user.User{}, *uc

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

func (s *UserService) UnregisterUser(ctx context.Context, userID string) []*user.UserError {

	uc := user.NewUserChecker()

	uc.UserID(userID)

	if uc.HasErrors() {
		return *uc
	}

	repo := repository.NewUserRepo()

	err := repo.Delete(ctx, userID)

	if err != nil {
		uc.AddError("db_error", "bad attempt on db user deletion", "db")
		return *uc
	}

	return nil

}

// [TODO] study the relationship between page offset and limit and see if there is a better handling for the
// values that the current implementation
func (s *UserService) ListUsers(ctx context.Context, params ListUsersParams) (user.UsersPage, []*user.UserError) {

	uc := user.NewUserChecker()

	var repoParams = repository.UserListParams{
		Limit: 10,
		Page:  1,
	}

	fmt.Println(params.Limit, params.Page)

	if params.Username != nil {
		repoParams.Username = *params.Username
	}

	if params.Role != nil {
		uc.Role(*params.Role)
		repoParams.Role = *params.Role
	}

	if params.Status != nil {
		uc.Status(*params.Status)
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

	if uc.HasErrors() {
		return user.UsersPage{}, *uc
	}

	repo := repository.NewUserRepo()

	if repoParams.Page == 1 {
		repoParams.Offset = 0
	} else {
		repoParams.Offset = int64(repoParams.Page) * repoParams.Limit
	}

	us, err := repo.List(ctx, repoParams)

	if err != nil {
		uc.AddError("db_error", "bad attempt on db user deletion", "db")
		return user.UsersPage{}, *uc
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

func (s *UserService) UpdateUser(ctx context.Context, userId string, roleID string, status string, email string, password string, username string) []*user.UserError {

	return []*user.UserError{}
}
