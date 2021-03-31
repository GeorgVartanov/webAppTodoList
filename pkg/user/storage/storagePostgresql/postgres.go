package storagePostgresql

import "github.com/GeorgVartanov/myWebApp/pkg/database/sqlPostgresDB"

type StorageUser struct {
	sqlPostgresDB.PostgresDB
}

func NewStorageUser(PDB sqlPostgresDB.PostgresDB) *StorageUser {
	return &StorageUser{PDB}
}
