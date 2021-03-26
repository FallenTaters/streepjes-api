package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/PotatoesFall/streepjes/shared/migrate"
)

var db *sql.DB

func main() {
	getDB()
	defer db.Close()

	_, err := db.Exec(string(MustAsset(`testdata.sql`)))
	if err != nil {
		panic(err)
	}
}

func getDB() {
	database, err := sql.Open("sqlite3", "./streepjes.db")
	if err != nil {
		panic(err)
	}
	db = database

	err = migrate.Migrate(db)
	if err != nil {
		panic(err)
	}
}
