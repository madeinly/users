package service

type RegisteUserParams struct {
	Username string
	Email    string
	Password string
	Role     string
	Status   string
}

type ListUsersParams struct {
	Username *string
	Role     *string
	Status   *string
	Page     *string
	Limit    *string
}

type UpdateUserParams struct {
	UserID   string
	Username string
	Role     string
	Status   string
	Email    string
	Password string
}
