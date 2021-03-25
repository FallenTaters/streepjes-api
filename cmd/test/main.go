package main

import (
	"database/sql"
	"fmt"

	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared/migrate"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	getDB()
	users.Init(db)
	err := users.Insert(`admin`, `admin`, `admin`, users.RoleAdmin)
	if err != nil {
		fmt.Println(err.Error())
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
