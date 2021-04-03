package create

import (
	"time"
)

//User ...
type User struct {
	//ID        int       `db:"user_id"`
	UserName  string    `json:"userName" db:"user_name"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  []byte    `json:"password" db:"password"`
	Created   time.Time `json:"created" db:"created"`
	Changed   time.Time `json:"changed" db:"changed"`
	IsAdmin   bool      `json:"isAdmin" db:"is_admin"`
}



func (u *User) GetQuery() string {
	return `INSERT INTO "users" ("user_name",
								"first_name",
								"last_name",
								"email",
								"password",
								"created",
								"changed",
								"is_admin")
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
}
