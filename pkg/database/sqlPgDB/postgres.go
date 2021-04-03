package sqlPgDB

import (
	"fmt"
	"log"

	"github.com/GeorgVartanov/myWebApp/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// PostgresDB Postgres class struct
type PostgresDB struct {
	*sqlx.DB
}

// NewPostgresDB get Postgres class struct
func NewPostgresDB(config config.DBPostgresConfig) *PostgresDB {
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	pg := PostgresDB{}
	if err := pg.connect(dbPath); err != nil {
		log.Fatalln(err)
	}
	return &pg
}

// connect connection to Postgresql
func (s *PostgresDB) connect(dbPath string) error {
	db, err := sqlx.Open("postgres", dbPath)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	db.SetMaxOpenConns(50)

	s.DB = db
	return nil
}
