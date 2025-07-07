package repository

import (
	"database/sql"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/user"
)

type Auth interface {
	ValidateCredentials(email string, password string) (bool, string)
}

type GetBy interface {
	GetByUsername(username string) user.User
	GetByEmail(email string) user.User
}

type Session interface {
	CreateSession()
	UpdateSession()
	ValidateSession()
	CheckExist()
}

type UserArgs struct {
	ID       string
	Username string
	Email    string
	Status   string
	Password string
	RoleID   string
}

type sqliteRepo struct {
	db *sql.DB
}

// var _ Repo = (*sqliteRepo)(nil)

func NewUserRepo() sqliteRepo {
	return sqliteRepo{db: core.DB()}
}
