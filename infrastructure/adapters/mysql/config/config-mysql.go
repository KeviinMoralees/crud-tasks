package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
}

func initDatabase() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
}
