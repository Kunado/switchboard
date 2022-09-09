package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=switchboard password=pass sslmode=disable")
	if err != nil {
		panic(err)
	}
}
