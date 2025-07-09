package repository

import (
	"database/sql"

	"github.com/madeinly/core"
)

type sqliteRepo struct {
	db *sql.DB
}

func NewUserRepo() sqliteRepo {
	return sqliteRepo{db: core.DB()}
}
