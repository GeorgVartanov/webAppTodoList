package pg

import (
	"github.com/GeorgVartanov/myWebApp/pkg/user/service/create"
)

func (c *UserStorage) Create(u *create.User) error {
	stmt, err := c.db.Preparex(u.GetQuery())
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		u.UserName,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.Created,
		u.Changed,
		u.IsAdmin,
	)
	if err != nil {
		return err
	}
	return nil
}
