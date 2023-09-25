package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Auth Auth
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Auth: NewAuthRepository(db),
	}
}
