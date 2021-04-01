package main

import (
	"database/sql"

	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
	"github.com/PotatoesFall/streepjes/shared/migrate"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	getDB()
	defer db.Close()
	users.Init(db)

	mustInsertUsers([]users.User{
		{
			Username: `adminG`,
			Club:     shared.ClubGladiators,
			Name:     `adminG`,
			Password: []byte(`admin`),
			Role:     users.RoleAdmin,
		}, {
			Username: `adminP`,
			Club:     shared.ClubParabool,
			Name:     `adminP`,
			Password: []byte(`admin`),
			Role:     users.RoleAdmin,
		}, {
			Username: `bar`,
			Name:     `bar`,
			Password: []byte(`bar`),
			Role:     users.RoleBartender,
		},
	})
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

func mustInsertUsers(u []users.User) {
	for _, user := range u {
		err := users.Insert(user)
		if err != nil {
			panic(err)
		}
	}
}
