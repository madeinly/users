package repository

type CreateUserParams struct {
	UserID   string
	Username string
	Email    string
	Password string
	Role     string
	Status   string
}

type UserListParams struct {
	Username string
	Status   string
	Role     string
	Offset   int64
	Page     int
	Limit    int64
}
