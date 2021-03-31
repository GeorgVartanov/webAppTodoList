package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/GeorgVartanov/myWebApp/config"
	"github.com/GeorgVartanov/myWebApp/pkg/database/sqlPostgresDB"
)

func main() {
	var config config.TomlConfig
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", config)

	pDB := sqlPostgresDB.NewPostgresDB(config.DBPostgres)
	defer pDB.Close()
}
