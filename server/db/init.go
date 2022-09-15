package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var Db *sql.DB

func init() {
	var err error
	connection := connectionStmt()
	Db, err = sql.Open("pgx", connection)
	if err != nil {
		panic(err)
	}
}

func connectionStmt() (statement string) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST not set. exit.")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Fatal("POSTGRES_USER not set. exit.")
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		log.Fatal("POSTGRES_PASSWORD not set. exit.")
	}
	statement = fmt.Sprintf("host=%s port=%s user=%s dbname=switchboard password=%s sslmode=disable", host, port, user, password)
	return
}
