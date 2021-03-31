package config

//TomlConfig
type TomlConfig struct {
	DBPostgres DBPostgresConfig `toml:"postgresql"`
}

//DBPostgresConfig
type DBPostgresConfig struct {
	Host     string //`toml:"host"`
	Port     string //`toml:"port"`
	User     string //`toml:"user"`
	Password string //`toml:"password"`
	DBName   string //`toml:"dbname"`
	SSLMode  string //`toml:"sslmode"`
}
