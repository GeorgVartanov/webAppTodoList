package main

import (
	"fmt"
	"github.com/GeorgVartanov/myWebApp/pkg/user/http/rest"
	"github.com/GeorgVartanov/myWebApp/pkg/user/storage/pg"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/GeorgVartanov/myWebApp/config"
	"github.com/GeorgVartanov/myWebApp/pkg/database/sqlPgDB"
)

func main() {
	var config config.TomlConfig
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	pDB := sqlPgDB.NewPostgresDB(config.DBPostgres)
	defer pDB.Close()

	newuser := rest.User{
		UserName:      "Zozo1",
		FirstName:     "Zozo1",
		LastName:      "zozo1",
		Email:         "zozo@",
		Password:      "zozozozo",
		PasswordCheck: "zozozozo",
		Created:       time.Now(),
		Changed:       time.Now(),
		IsAdmin:       false,
	}
	err := newuser.ValidateFields()
	if err != nil {
		fmt.Println(err)
	}
	userService, err := newuser.ConvertToServiceUser()

	UserStorage := pg.NewUserStorage(pDB)
	err = UserStorage.Create(userService)
	if err != nil {
		fmt.Println(err)
	}
}

//migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path pkg/database/sqlPgDB/migrations up

//migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path pkg/database/sqlPgDB/migrations up
