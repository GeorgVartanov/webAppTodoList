package pg

import (
	"github.com/GeorgVartanov/myWebApp/pkg/user/service/adding"
)

func (c *UserStorage) AddUser(u *adding.User) error {
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
