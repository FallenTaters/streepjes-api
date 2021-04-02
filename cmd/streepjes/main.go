package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"go.etcd.io/bbolt"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/catalog"
	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/orders"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
	"github.com/PotatoesFall/streepjes/shared/migrate"
)

var db *bbolt.DB

const path = "streepjes.db"

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

	au := r.Group(`/`, authMiddleware)
	au.POST(`/logout`, postLogout)
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
	database, err := bbolt.Open(path, 0666, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	db = database

	err = migrate.Migrate(database)
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
