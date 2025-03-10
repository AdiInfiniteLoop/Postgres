package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func DBConnection() error {
	var err error
	DB, err = sql.Open("postgres", "postgres://adiinfiniteloop:mysecretpassword@localhost:5432/mydatabase")

	if err != nil {
		return err
	}
	return nil
}

func DBClose() error {
	return DB.Close()
}
