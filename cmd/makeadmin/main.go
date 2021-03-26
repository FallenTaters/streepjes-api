package main

import (
	"database/sql"

	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared/migrate"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	getDB()
	defer db.Close()
	users.Init(db)

	err := users.Insert(`admin`, `admin`, `admin`, users.RoleAdmin)
	if err != nil {
		panic(err)
	}

	err = users.Insert(`bartender`, `bartender`, `bartender`, users.RoleBartender)
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
