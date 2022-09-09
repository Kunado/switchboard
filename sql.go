package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		log.Fatalf("%v", err)
	}
	//dbName := "switchboard"
	//_, err = db.Exec("create database " + dbName)
	//if err != nil {
	//	log.Fatalf("%v", err)
	//}

	_, err = db.Exec("CREATE TABLE test ( id integer, username varchar(255) )")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
