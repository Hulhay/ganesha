package config

import (
	"os"
	"strconv"

	"github.com/hulhay/ganesha/lib/database"
)

type Config struct {
	port   int
	dbSqlx database.SqlxDatabase
}

func NewConfig() *Config {
	cfg := new(Config)

	cfg.ConnectSqlxDB()

	return cfg
}

func (c *Config) Port() int {
	v := os.Getenv("PORT")
	c.port, _ = strconv.Atoi(v)

	return c.port
}

func (c *Config) ConnectSqlxDB() {
	c.dbSqlx = database.InitSQLX()
}

func (c *Config) DisconnectDB() {
	c.dbSqlx.Close()
}
