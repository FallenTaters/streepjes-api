package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/user"
	"github.com/PotatoesFall/streepjes/shared"
)

func main() {
	readSettings()
	initStuff()

	r := router.New()

	r.ErrorHandler = errorHandler
	r.Reader = reader

	r.POST(`/login`, login)

	a := r.Group(`/`, authMiddleware)
	a.GET(`/order`, getOrder)

	panic(r.Start(`:` + settings.Port))
}

func initStuff() {
	shared.Init(settings.DisableSecure)
}

func errorHandler(c *router.Context, v interface{}) {
	fmt.Printf("panic: %#v\n", v)
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

func login(c *router.Context, credentials user.Credentials) error {
	role := user.LogIn(c.Response, credentials)
	if role == user.RoleNotAuthorized {
		return c.String(http.StatusBadRequest, `invalid username or password`)
	}

	shared.SetCookie(c.Response, user.AuthCookie, `token`, 5*60)
	return c.JSON(http.StatusOK, role)
}

func getOrder(c *router.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
