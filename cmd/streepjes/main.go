package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
	"github.com/PotatoesFall/streepjes/shared/buckets"
)

func main() {
	readSettings()

	shared.Init(settings.DisableSecure)
	close := buckets.Init()
	defer close()

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
