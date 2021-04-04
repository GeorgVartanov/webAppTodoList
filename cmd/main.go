package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/GeorgVartanov/myWebApp/config"
	"github.com/GeorgVartanov/myWebApp/pkg/database/sqlPgDB"
	"github.com/GeorgVartanov/myWebApp/pkg/user/http/rest"
	"github.com/GeorgVartanov/myWebApp/pkg/user/service/adding"
	"github.com/GeorgVartanov/myWebApp/pkg/user/storage/pg"
	"github.com/gin-gonic/gin"
)

func main() {
	var config config.TomlConfig
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	pDB := sqlPgDB.NewPostgresDB(config.DBPostgres)
	defer pDB.Close()
	UserStorage := pg.NewUserStorage(pDB)
	adser := adding.NewService(UserStorage)

	r := gin.Default()
	api := r.Group("/api/user")

	rest.Handler(api, adser)
	r.Run(":3000")

}

//migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path pkg/database/sqlPgDB/migrations up

//migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path pkg/database/sqlPgDB/migrations up
