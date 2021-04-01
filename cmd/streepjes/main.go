package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/catalog"
	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/orders"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
	"github.com/PotatoesFall/streepjes/shared/migrate"
)

var db *sql.DB

func main() {
	readSettings()
	getDB()
	defer db.Close()
	initStuff()

	r := router.New()

	r.Use(corsMiddleware)
	r.ErrorHandler = errorHandler
	r.Reader = reader

	r.POST(`/login`, postLogin)
	r.POST(`/logout`, postLogout)

	au := r.Group(`/`, authMiddleware)
	au.POST(`/active`, postActive)
	au.GET(`/catalog`, getCatalog)
	au.GET(`/members`, getMembers)
	au.POST(`/order`, postOrder)
	au.GET(`/orders`, getOrders)

	ad := au.Group(`/`, roleMiddleware(users.RoleAdmin))
	ad.GET(`/users`, getUsers)

	panic(r.Start(`:` + settings.Port))
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

func initStuff() {
	shared.Init(settings.DisableSecure)
	catalog.Init(db)
	users.Init(db)
	members.Init(db)
	orders.Init(db)
}

func errorHandler(c *router.Context, v interface{}) {
	fmt.Fprintf(os.Stderr, "panic: %#v\n", v)
	_ = c.NoContent(http.StatusInternalServerError)
}

func reader(c *router.Context, dst interface{}) (bool, error) {
	defer c.Request.Body.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return false, c.String(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(body, dst)
	if err != nil {
		return false, c.String(http.StatusBadRequest, err.Error())
	}

	return true, nil
}
