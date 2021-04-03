package pg

import (
	"github.com/GeorgVartanov/myWebApp/pkg/database/sqlPgDB"
)

type UserStorage struct {
	db *sqlPgDB.PostgresDB
}

func NewUserStorage(PDB *sqlPgDB.PostgresDB) *UserStorage {
	return &UserStorage{PDB}
}
