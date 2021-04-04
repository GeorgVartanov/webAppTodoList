package deleting

type User struct {
	UserName  string    `json:"userName" db:"user_name"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
}
