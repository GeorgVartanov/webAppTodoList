package storagePostgresql

import "github.com/GeorgVartanov/myWebApp/pkg/user/storage"

func (u *StorageUser) CreateUserRepository(user *storage.User) error {

	tx := u.MustBegin()

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	_, err := tx.NamedExec(`INSERT INTO "myUser" ("firstName", "lastName", "email", "password", "created", "changed", "isAdmin") VALUES (:firstName, :lastName, :email, :password, :created, :changed, :isAdmin) RETURNING id`, &user)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	//fmt.Println(id)
	return nil
}
