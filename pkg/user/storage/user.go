package storage

import "time"

//User ...
type User struct {
	ID        int       `db:"user_id"`
	UserName  string    `db:"user_name"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Created   time.Time `db:"created"`
	Changed   time.Time `db:"changed"`
	IsAdmin   bool      `db:"is_admin"`
}

// postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
//migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path pkg/database/sqlPostgresDB/migrations up
