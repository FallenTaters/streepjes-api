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
	defer close() //nolint: errcheck

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
	au.POST(`/order/delete/:id`, postOrderDelete)
	au.GET(`/club`, getClub)

	ad := au.Group(`/`, roleMiddleware(users.RoleAdmin))
	ad.GET(`/users`, getUsers)
	ad.POST(`/user`, postUser)
	ad.POST(`/user/delete/:username`, postUserDelete)

	ad.POST(`/member`, postMember)
	ad.POST(`/member/delete/:id`, postMemberDelete)

	ad.POST(`/category`, postCategory)
	ad.POST(`/category/delete/:id`, postCategoryDelete)
	ad.POST(`/product`, postProduct)
	ad.POST(`/product/delete/:id`, postProductDelete)

	ad.GET(`/orders/:year/:month`, getOrdersByMonth)

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
