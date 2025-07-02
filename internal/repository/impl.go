package repository

import (
	"database/sql"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/user"
)

type Core interface {
	Create(args UserArgs) (string, error)
	Delete(userID string) error
	Update(args UserArgs) error
	GetByID(userID string) user.User
	List(args UserListArgs) []user.User
}

type Auth interface {
	ValidateCredentials(email string, password string) (bool, string)
}

type GetBy interface {
	GetByUsername(username string) user.User
	GetByEmail(email string) user.User
}

type Repo interface {
	Core
	Auth
	GetBy
}

type UserArgs struct {
	ID       string
	Username string
	Email    string
	Status   string
	Password string
	RoleID   string
}

type UserListArgs struct {
	Username string
	Status   string
	RoleID   int64
	Offset   int64
	Page     int
	Limit    int64
}

type sqliteRepo struct {
	db *sql.DB
}

var _ Repo = (*sqliteRepo)(nil)

func NewUserRepo() sqliteRepo {
	return sqliteRepo{db: core.DB()}
}
