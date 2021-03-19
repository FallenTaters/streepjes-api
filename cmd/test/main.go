package main

import (
	"database/sql"

	"github.com/PotatoesFall/streepjes/domain/user"
	"github.com/PotatoesFall/streepjes/shared/migrate"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	getDB()
	user.Init(db)
	panic(user.Insert(`admin`, `admin`, `admin`, user.RoleAdmin))
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
